package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	config "prtvi/expense-tracker/config"
	model "prtvi/expense-tracker/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BsonDocToSummary(doc bson.M) model.Summary {
	var summary model.Summary
	docByte, _ := json.Marshal(doc)
	json.Unmarshal(docByte, &summary)

	return summary
}

func initSummaryObj() model.Summary {
	var summary model.Summary

	indi := make(model.IndiModeSums, len(config.AllModesOfPayment))

	for key := range config.AllModesOfPayment {
		indi[key] = model.Mode{}
	}
	summary.IndiModeSums = indi

	return summary
}

// insert the initialized summary doc in db
func InitSummary() error {
	summary := initSummaryObj()

	// create filter, update and options for querying
	filter := bson.M{}
	update := bson.M{
		"$set": bson.D{
			{Key: "total_income", Value: summary.TotalIncome},
			{Key: "total_expense", Value: summary.TotalExpense},
			{Key: "total_balance", Value: summary.TotalBalance},
			{Key: "indi_mode_sums", Value: summary.IndiModeSums},
		},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	cursor := config.Summary.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if cursor.Err() != nil {
		return cursor.Err()
	}

	return nil
}

func extractTransactionsForMode(mode string, startDate, endDate time.Time, sort string) []model.Transaction {
	findOptions := getSortOptions(sort)

	// create filter to get transactions within given range but of the current "mode"
	filter := bson.M{config.DateID: bson.M{"$gte": startDate, "$lte": endDate}, "mode": mode}

	// run query
	cursor, err := config.Transactions.Find(context.TODO(), filter, findOptions)
	if err != nil {
		fmt.Println(err.Error())
	}

	// extract results
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())
	}

	// convert to obj
	extractedTs := make([]model.Transaction, len(results))
	for i, resultItem := range results {
		extractedTs[i] = BsonDocToTransaction(resultItem)
	}

	return extractedTs
}

// return a model.Summary object for the transactions found in the time range
func GetSummary(startDate, endDate time.Time, sort string) model.Summary {
	// init summary obj
	summary := initSummaryObj()

	// get the modes of payment opted
	_, mop := GetCurrencyAndModesOfPayment()

	// loop over all modes of payment and make query for transactions for each mode type
	for mode, value := range mop {
		// if mop is not opted then skip iteration
		if !value.IsChecked {
			continue
		}

		extractedTs := extractTransactionsForMode(mode, startDate, endDate, sort)

		// init modeSum for the current "mode"
		indiModeSum := model.Mode{}
		for _, t := range extractedTs {
			if t.Type == config.TypeIncomeID {
				indiModeSum.Income += t.Amount
			} else {
				indiModeSum.Expense += t.Amount
			}
		}

		// attach to summary object
		summary.IndiModeSums[mode] = indiModeSum
	}

	// calculate total income, expense and current balance
	for _, value := range summary.IndiModeSums {
		summary.TotalIncome += value.Income
		summary.TotalExpense += value.Expense
	}

	summary.TotalBalance = summary.TotalIncome - summary.TotalExpense

	return summary
}

// updates the summary into the db
func UpdateMainSummary() (model.Summary, time.Time, time.Time) {
	_, year := GetCurrentMonthAndYear()
	startDate, endDate, _ := GetNewestAndOldestTDates(year)
	summary := GetSummary(startDate, endDate, config.SortAscID)

	// create filter, update and options for querying
	filter := bson.M{}
	update := bson.M{
		"$set": bson.D{
			{Key: "total_income", Value: summary.TotalIncome},
			{Key: "total_expense", Value: summary.TotalExpense},
			{Key: "total_balance", Value: summary.TotalBalance},
			{Key: "indi_mode_sums", Value: summary.IndiModeSums},
		},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	cursor := config.Summary.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if cursor.Err() != nil {
		fmt.Println("err")
	}

	return summary, startDate, endDate
}

func FetchMainSummary() model.Summary {
	cursor := config.Summary.FindOne(context.TODO(), bson.M{})

	fetchedDoc := bson.M{}
	decodeErr := cursor.Decode(&fetchedDoc)
	if decodeErr != nil {
		fmt.Println("error")
	}

	return BsonDocToSummary(fetchedDoc)
}
