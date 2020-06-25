package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CloudWatchEvent) error {
	fmt.Printf("event: %s", event)
	sm := &SlackMessage{Text: "test message"}
	return sm.post()
}

func main() {
	lambda.Start(handler)
}
