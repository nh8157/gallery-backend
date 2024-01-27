package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nh8157/gallery-backend/pkg/handler"
)

func main() {
	lambda.Start(handler.HandleRequests)
}
