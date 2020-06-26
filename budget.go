package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/sts"
)

var budgetName = os.Getenv("BUDGET_NAME")

func createBudgetMessage() (string, error) {
	sess := session.Must(session.NewSession())
	callerIdentity, err := sts.New(sess).GetCallerIdentity(nil)
	if err != nil {
		return "", fmt.Errorf("failed GetCallerIdentity: %s", err)
	}
	accountId := callerIdentity.Account

	input := &budgets.DescribeBudgetInput{
		AccountId:  accountId,
		BudgetName: &budgetName,
	}
	output, err := budgets.New(sess).DescribeBudget(input)
	if err != nil {
		return "", fmt.Errorf("failed DescribeBudget: %s", err)
	}

	spend := *output.Budget.CalculatedSpend
	return fmt.Sprintf("Actual: %s USD, Forecasted %s USD",
		*spend.ActualSpend.Amount, *spend.ForecastedSpend.Amount), nil
}
