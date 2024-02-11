package dynamo

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/nh8157/gallery-backend/internal/metadata"
)

var svc *dynamodb.DynamoDB

func init() {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		log.Fatal(err)
	}
	svc = dynamodb.New(sess)
}

func DynamoPut(tableName string, items map[string]*dynamodb.AttributeValue) error {
	_, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      items,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DynamoGetItem(tableName string, fileName string) (*metadata.Metadata, error) {
	data, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("gallery"),
		Key: map[string]*dynamodb.AttributeValue{
			"FileName": {
				S: aws.String(fileName),
			},
		},
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var md metadata.Metadata
	err = dynamodbattribute.UnmarshalMap(data.Item, &md)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &md, nil
}

// could be extended to require value of certain fields
func DynamoGetItems(tableName string) ([]metadata.Metadata, error) {
	proj := expression.NamesList(
		expression.Name(string(metadata.FileName)),
		expression.Name(string(metadata.Model)),
		expression.Name(string(metadata.LensModel)),
		expression.Name(string(metadata.FocalLength)),
		expression.Name(string(metadata.DateTime)),
		expression.Name(string(metadata.ApertureValue)),
		expression.Name(string(metadata.ISOSpeedRatings)),
		expression.Name(string(metadata.ShutterSpeedValue)),
	)

	// expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	data, err := svc.Scan(&dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	mds := []metadata.Metadata{}
	for _, item := range data.Items {
		var md metadata.Metadata
		err = dynamodbattribute.UnmarshalMap(item, &md)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println(md.FileName)
		mds = append(mds, md)
	}
	return mds, nil
}

func DynamoDelete(tableName string, key string) error {
	_, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"FileName": {
				S: aws.String(key),
			},
		},
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DynamoPatch(tableName string) {

}
