package handler

import "github.com/aws/aws-lambda-go/events"

func BuildResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		Body:       body,
	}
}
