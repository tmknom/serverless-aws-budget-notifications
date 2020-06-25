package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CloudWatchEvent) error {
	fmt.Printf("event: %s", event)
	budget, err := describeBudget()
	if err != nil {
		return err
	}

	text := fmt.Sprintf("Actual: %s USD, Forecasted: %s USD", budget.Actual, budget.Forecasted)
	sm := &SlackMessage{Text: text}
	return sm.post()
}

func main() {
	lambda.Start(handler)
}
