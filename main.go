package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
)

var incomingWebhookURL = os.Getenv("INCOMING_WEBHOOK_URL")

func handler() error {
	message, err := createBudgetMessage()
	if err != nil {
		return err
	}

	webhookMessage := &slack.WebhookMessage{Text: message}
	return slack.PostWebhook(incomingWebhookURL, webhookMessage)
}

func main() {
	lambda.Start(handler)
}
