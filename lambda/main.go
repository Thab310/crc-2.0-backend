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

	if tableName == "" || region == "" {
		return 0, fmt.Errorf("TABLE_NAME and REGION environment variables must be set")
	}

	// Initialize a session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess, aws.NewConfig().WithRegion(region))

	// Get the item with ID "0"
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

	if result.Item == nil {
		// Item does not exist, initialize it
		item = Item{
			ID:    "0",
			Views: 0,
		}
	} else {
		// Item exists, unmarshal it
		err = dynamodbattribute.UnmarshalMap(result.Item, &item)
		if err != nil {
			return 0, fmt.Errorf("Failed to unmarshal Record, %v", err)
		}

		// Increment views
		item.Views++
	}

	// Print the current views (for debugging)
	fmt.Println("Current views:", item.Views)

	// Marshal the updated item
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return 0, fmt.Errorf("Got error marshalling new item: %s", err)
	}

	// Put the updated item back
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
