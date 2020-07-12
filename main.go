package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kylerwsm/kyler-bot/pkg/models"
	"github.com/kylerwsm/kyler-bot/pkg/util"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (models.Response, error) {
	var update models.Update
	err := json.Unmarshal([]byte(request.Body), &update)
	if err == nil {
		util.SendToUser(update.Message.From.ID, "Hello")
		return models.Response{StatusCode: http.StatusOK}, nil
	}
	return models.Response{StatusCode: http.StatusInternalServerError}, err
}

func main() {
	lambda.Start(Handler)
}
