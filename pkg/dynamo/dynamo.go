package dynamo

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

func DynamoGet(talbeName string) {
}

func DynamoDelete(tableName string) {

}

func DynamoPatch(tableName string) {

}
