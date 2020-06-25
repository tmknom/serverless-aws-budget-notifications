package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/sts"
)

var budgetName = os.Getenv("BUDGET_NAME")

type Budget struct {
	Actual     string
	Forecasted string
}

func describeBudget() (*Budget, error) {
	sess := session.Must(session.NewSession())
	callerIdentity, err := sts.New(sess).GetCallerIdentity(nil)
	if err != nil {
		return nil, fmt.Errorf("failed GetCallerIdentity: %s", err)
	}

	input := &budgets.DescribeBudgetInput{
		AccountId:  callerIdentity.Account,
		BudgetName: &budgetName,
	}
	output, err := budgets.New(sess).DescribeBudget(input)
	if err != nil {
		return nil, fmt.Errorf("failed DescribeBudget: %s, input: %s", err, input)
	}

	return &Budget{
		Actual:     *output.Budget.CalculatedSpend.ActualSpend.Amount,
		Forecasted: *output.Budget.CalculatedSpend.ForecastedSpend.Amount,
	}, nil
}
