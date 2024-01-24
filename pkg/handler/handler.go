package handler

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleHealthCheck(requests events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := BuildResponse(200, "")
	return response, nil
}

func HandleGallery(requests events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	var err error
	switch requests.HTTPMethod {
	case "GET":
		response, err = HandleGalleryGet(requests)
	case "POST":
		response, err = HandleGalleryPost(requests)
	case "PATCH":
		response, err = HandleGalleryPatch(requests)
	case "DELETE":
		response, err = HandleGalleryDelete(requests)
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
