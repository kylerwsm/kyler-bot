package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kylerwsm/kyler-bot/pkg/entity"
	"github.com/kylerwsm/kyler-bot/pkg/services"
)

// Handler is our lambda handler.
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (entity.Response, error) {
	var update entity.Update
	err := json.Unmarshal([]byte(request.Body), &update)
	if err != nil {
		return entity.Response{StatusCode: http.StatusInternalServerError}, err
	}

	err = services.AddFoodItem(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		return entity.Response{StatusCode: http.StatusInternalServerError}, err
	}

	foodItems, err := services.GetFoodItems(update.Message.Chat.ID)
	if err != nil {
		return entity.Response{StatusCode: http.StatusInternalServerError}, err
	}

	err = services.SendToUser(update.Message.Chat.ID, strings.Join(foodItems, "\n"))
	if err != nil {
		return entity.Response{StatusCode: http.StatusInternalServerError}, err
	}
	return entity.Response{StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(Handler)
}
