package models

import "github.com/aws/aws-lambda-go/events"

// Response type is APIGatewayProxyResponse to leverage AWS Lambda Proxy Request functionality.
type Response events.APIGatewayProxyResponse

// Update is a Telegram object that the handler receives every time an user interacts with the bot.
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message indicates details found in an update.
type Message struct {
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

// From indicates the message sender.
type From struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

// Chat indicates the relevant conversation.
type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Type      string `json:"type"`
}
