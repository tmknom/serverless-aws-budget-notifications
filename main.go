package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CloudWatchEvent) {
	fmt.Printf("event: %s", event)
	fmt.Printf("env.SLACK_INCOMING_WEBHOOK_URL: %s", os.Getenv("SLACK_INCOMING_WEBHOOK_URL"))
}

func main() {
	lambda.Start(handler)
}
