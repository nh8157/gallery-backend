package handler

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func HandleHealthCheck(requests events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := BuildResponse(200, "")
	return response, nil
}

func HandleGallery(requests events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	var err error

	sess, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("aws_region"))})
	if err != nil {
		err := fmt.Errorf("Failed to establish connection with aws")
		response := BuildResponse(502, err.Error())
		return response, err
	}

	s3 := s3.New(sess)

	switch requests.HTTPMethod {
	case "GET":
		response, err = HandleGalleryGet(requests, s3)
	case "PUT":
		response, err = HandleGalleryPut(requests, s3)
	case "PATCH":
		response, err = HandleGalleryPatch(requests, s3)
	case "DELETE":
		response, err = HandleGalleryDelete(requests, s3)
	}
	return response, err
}

func HandleRequests(requests events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	var err error
	switch requests.Path {
	case "/health":
		response, err = HandleHealthCheck(requests)
	case "/gallery":
		response, err = HandleGallery(requests)
	default:
		err = fmt.Errorf("%s is not a valid path", requests.Path)
		response = BuildResponse(404, err.Error())
	}
	if err != nil {
		return response, err
	}
	return response, nil
}
