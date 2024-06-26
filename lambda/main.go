package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	ID    string `json:"ID"`
	Views int    `json:"Views"`
}

func handleRequest() (int, error) {
	// Get environment variables
	tableName := os.Getenv("TABLE_NAME")
	region := os.Getenv("REGION")

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess, aws.NewConfig().WithRegion(region))

	// Get the item with id "0"
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String("0"),
			},
		},
	})

	if err != nil {
		return 0, fmt.Errorf("Got error calling GetItem: %s", err)
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return 0, fmt.Errorf("Failed to unmarshal Record, %v", err)
	}

	// Increment views
	item.Views++

	fmt.Println(item.Views)

	// Put the updated item back
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return 0, fmt.Errorf("Got error marshalling new item: %s", err)
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	})

	if err != nil {
		return 0, fmt.Errorf("Got error calling PutItem: %s", err)
	}

	return item.Views, nil
}

func main() {
	lambda.Start(handleRequest)
}
