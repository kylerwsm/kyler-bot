package repositories

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/kylerwsm/kyler-bot/pkg/entity"
)

// Declare DynamoDB instance which is safe for concurrent use.
var svc = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-1"))

const tableName = "FoodList"

// AddFoodItem adds a food item in the database.
func AddFoodItem(chatID int, foodPlace string) error {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	foodList := entity.FoodPlaceEntry{ChatID: chatID, Timestamp: timestamp, FoodItem: foodPlace}
	av, err := dynamodbattribute.MarshalMap(foodList)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = svc.PutItem(input)
	return err
}

// RemoveFoodItem removes a food item in the database.
func RemoveFoodItem(chatID int, foodItem string) error {
	return nil
}

// GetFoodItems returns the chat's food list in the database.
func GetFoodItems(chatID int) ([]string, error) {
	result, err := svc.Scan(&dynamodb.ScanInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":cid": {
				N: aws.String(strconv.Itoa(chatID)),
			},
		},
		FilterExpression: aws.String("ChatID = :cid"),
		TableName:        aws.String(tableName),
	})
	if err != nil {
		return []string{}, err
	}
	var foodPlaces []string
	for _, item := range result.Items {
		var foodPlaceEntry entity.FoodPlaceEntry
		err = dynamodbattribute.UnmarshalMap(item, &foodPlaceEntry)
		if err != nil {
			return []string{}, err
		}
		foodPlaces = append(foodPlaces, foodPlaceEntry.FoodItem)
	}
	return foodPlaces, nil
}

// ClearFoodItems clears the chat's food list in the database.
func ClearFoodItems(chatID int) error {
	return nil
}
