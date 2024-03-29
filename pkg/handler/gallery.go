package handler

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nh8157/gallery-backend/internal/metadata"
	"github.com/nh8157/gallery-backend/internal/response"
	"github.com/nh8157/gallery-backend/pkg/dynamo"
)

func HandleGalleryGet(requests events.APIGatewayProxyRequest, svc *s3.S3) (events.APIGatewayProxyResponse, error) {
	var errResp response.ErrorResponse
	// retrieve metadata
	mds, err := dynamo.DynamoGetItems("gallery")
	if err != nil {
		log.Printf("unable to get items: %s\n", err)
		errResp.Msg = err.Error()
		return BuildResponse(502, response.ToString(&errResp)), nil
	}
	// get presigned urls from s3
	for i := range mds {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("s3_bucket_name")),
			Key:    aws.String(*mds[i].FileName),
		})
		urlStr, err := req.Presign(60 * time.Minute)
		if err != nil {
			log.Printf("unable to generate presigned url for %s", *mds[i].FileName)
			errResp.Msg = response.S3Error
			return BuildResponse(502, response.ToString(&errResp)), nil
		}
		mds[i].S3Url = &urlStr
	}
	marshalled, err := json.Marshal(mds)
	if err != nil {
		log.Printf("unable to marshal JSON: %s\n", err)
		errResp.Msg = err.Error()
		return BuildResponse(500, response.ToString(&errResp)), err
	}
	return BuildResponse(200, string(marshalled)), nil
}

func HandleGalleryPut(requests events.APIGatewayProxyRequest, svc *s3.S3) (events.APIGatewayProxyResponse, error) {
	var md metadata.Metadata
	var errResp response.ErrorResponse
	err := json.Unmarshal([]byte(requests.Body), &md)
	if err != nil {
		log.Printf("cannot marshal JSON: %s\n", err)
		errResp.Msg = response.JsonParseError
		response := BuildResponse(500, response.ToString(&errResp))
		return response, nil
	}

	item, err := dynamo.DynamoGetItem("gallery", *md.FileName)
	if err != nil {
		log.Printf("unable to get items: %s\n", err)
		errResp.Msg = response.DynamoDbError
		return BuildResponse(502, response.ToString(&errResp)), nil
	} else if !item.IsNil() {
		log.Printf("file duplication: %s already exists\n", metadata.FileName)
		errResp.Msg = response.FileDuplicationError
		return BuildResponse(400, response.ToString(&errResp)), nil
	}

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("s3_bucket_name")),
		Key:    aws.String(*md.FileName),
	})
	urlStr, err := req.Presign(5 * time.Minute)

	if err != nil {
		log.Printf("unable to presign url: %s\n", err)
		errResp.Msg = response.S3Error
		response := BuildResponse(502, response.ToString(&errResp))
		return response, nil
	}

	jsonResponse := map[string]string{"url": urlStr}
	marshalledResponse, err := json.Marshal(jsonResponse)
	if err != nil {
		log.Printf("unable to marshal JSON: %s\n", err)
		errResp.Msg = response.JsonParseError
		response := BuildResponse(500, response.ToString(&errResp))
		return response, nil
	}

	response := BuildResponse(200, string(marshalledResponse))
	return response, nil
}

func HandleGalleryPatch(requests events.APIGatewayProxyRequest, svc *s3.S3) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	return response, nil
}

func HandleGalleryDelete(requests events.APIGatewayProxyRequest, svc *s3.S3) (events.APIGatewayProxyResponse, error) {
	var md metadata.Metadata
	var errResp response.ErrorResponse
	err := json.Unmarshal([]byte(requests.Body), &md)
	if err != nil {
		log.Printf("unable to unmarshal JSON: %s\n", err)
		errResp.Msg = response.JsonParseError
		response := BuildResponse(500, response.ToString(&errResp))
		return response, nil
	}

	err = dynamo.DynamoDelete("gallery", *md.FileName)
	if err != nil {
		log.Printf("unable to delete metadata from database: %s\n", err)
		errResp.Msg = response.DynamoDbError
		response := BuildResponse(502, response.ToString(&errResp))
		return response, nil
	}

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("s3_bucket_name")),
		Key:    aws.String(*md.FileName),
	})

	if err != nil {
		log.Printf("unable to delete object from S3 bucket: %s\n", err)
		errResp.Msg = response.S3Error
		response := BuildResponse(502, response.ToString(&errResp))
		return response, nil
	}

	response := BuildResponse(200, "")
	return response, nil
}
