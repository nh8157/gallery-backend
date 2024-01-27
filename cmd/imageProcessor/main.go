package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nh8157/gallery-backend/pkg/dynamo"
	"github.com/nh8157/gallery-backend/pkg/metadataReader"
)

func main() {
	lambda.Start(handleUploadEvent)
}

func handleUploadEvent(event events.S3Event) {
	bucket := event.Records[0].S3.Bucket.Name
	key := event.Records[0].S3.Object.Key

	sess, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("aws_region"))})
	if err != nil {
		log.Println(err)
		return
	}
	svc := s3.New(sess)

	output, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Println(err)
		return
	}

	md := metadataReader.ReadExif(output.Body, &key)
	// write the metadata to dynamoDB
	marshalled, err := dynamodbattribute.MarshalMap(md.ToMap())
	if err != nil {
		log.Println(err)
		return
	}

	err = dynamo.DynamoPut("gallery", marshalled)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Saved to DynamoDB")
}
