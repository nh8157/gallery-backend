package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func HandleGalleryGet(requests events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	var err error

	return response, err
}

func HandleGalleryPost(requests events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("aws_region"))})
	if err != nil {
		err := fmt.Errorf("Failed to establish connection with aws")
		response := BuildResponse(502, err.Error())
		return response, err
	}

	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("s3_bucket_name")),
		Key:    aws.String("myKey"),
	})
	urlStr, err := req.Presign(5 * time.Minute)

	if err != nil {
		err := fmt.Errorf("Failed to sign request")
		response := BuildResponse(502, err.Error())
		return response, err
	}

	response := BuildResponse(200, urlStr)
	return response, nil
}

func HandleGalleryPatch(requests events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	var err error
	return response, err
}

func HandleGalleryDelete(requests events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	var err error
	return response, err
}
