package handler

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Metadata struct {
	FileName string `json:"fileName"`
}

func HandleGalleryGet(requests events.APIGatewayProxyRequest, svc *s3.S3) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	var err error
	// query dynamoDB and return a list of s3 bucket urls
	return response, err
}

func HandleGalleryPut(requests events.APIGatewayProxyRequest, svc *s3.S3) (events.APIGatewayProxyResponse, error) {
	var metadata Metadata
	json.Unmarshal([]byte(requests.Body), &metadata)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("s3_bucket_name")),
		Key:    aws.String(metadata.FileName),
	})
	urlStr, err := req.Presign(5 * time.Minute)

	if err != nil {
		log.Println(err)
		response := BuildResponse(502, err.Error())
		return response, err
	}

	jsonResponse := map[string]string{"url": urlStr}
	marshalledResponse, err := json.Marshal(jsonResponse)
	if err != nil {
		log.Println(err)
		response := BuildResponse(500, err.Error())
		return response, err
	}

	response := BuildResponse(200, string(marshalledResponse))
	return response, nil
}

func HandleGalleryPatch(requests events.APIGatewayProxyRequest, svc *s3.S3) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	var err error
	return response, err
}

func HandleGalleryDelete(requests events.APIGatewayProxyRequest, svc *s3.S3) (events.APIGatewayProxyResponse, error) {
	var err error
	var metadata Metadata
	json.Unmarshal([]byte(requests.Body), &metadata)

	// search if the object is present in dynamoDB table

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("s3_bucket_name")),
		Key:    aws.String(metadata.FileName),
	})

	if err != nil {
		log.Println(err)
		response := BuildResponse(502, err.Error())
		return response, err
	}

	response := BuildResponse(200, "")
	return response, nil
}
