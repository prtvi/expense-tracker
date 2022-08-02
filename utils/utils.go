package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	config "prtvi/expense-tracker/config"
	model "prtvi/expense-tracker/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// convert primitive.M to model.Transaction object

func BsonDocToTransaction(doc bson.M) model.Transaction {
	var transaction model.Transaction
	docByte, _ := json.Marshal(doc)
	json.Unmarshal(docByte, &transaction)

	return transaction
}

func BsonDocToSummary(doc bson.M) model.Summary {
	var summary model.Summary
	docByte, _ := json.Marshal(doc)
	json.Unmarshal(docByte, &summary)

	return summary
}

// constructor for generating model.ResponseMessage objects

func CreateResponseMessage(statusCode int, success bool, message string) model.ResponseMsg {
	return model.ResponseMsg{StatusCode: statusCode, Success: success, Message: message}
}

func GetClassNameByValue(value float32) string {
	if value >= 0 {
		return config.ClassTTypeIncome
	}
	return config.ClassTTypeExpense
}

// date object to date string with format: 2022-05-25
func FormatDateShort(d time.Time) string {
	return d.String()[0:config.FORMAT_DATE_STR_LEN]
}

// date object to date string with format: Wed, 25 May 2022
func FormatDateLong(d time.Time) string {
	return d.Format(time.RFC1123Z)[0:config.FORMAT_DATE_STR_LEN_LONG]
}

// converts a single transaction from model.Transaction to model.TransactionFormatted (not for view, to be loaded on t-form)

func FormatTransaction(t model.Transaction) model.TransactionFormatted {
	// date object to format: 2022-05-25
	return model.TransactionFormatted{
		ID:     t.ID,
		Date:   FormatDateShort(t.Date),
		Desc:   t.Desc,
		Amount: t.Amount,
		Mode:   t.Mode,
		Type:   t.Type,
		PaidTo: t.PaidTo,
	}
}

// formats a single transaction for view, model.Traansaction to model.TransactionFormatted (format for view)
// truncate desc text to MAX_DESC_LEN & format date to format: Wed, 25 May 2022
// to view only on table and modal

func FormatTransactionForView(t model.Transaction) model.TransactionFormatted {
	T := FormatTransaction(t)

	// 1. format date to format into: Wed, 25 May 2022
	T.Date = FormatDateLong(t.Date)

	// 2. truncating desc text
	if len(t.Desc) > config.MAX_DESC_LEN {
		T.Desc = t.Desc[0:config.MAX_DESC_LEN] + "..."
	}

	return T
}

// specifically for "/get" route, to format an array of model.Transaction to array of model.TransactionFormatted, format for view

func FormatTransactionsForView(allTransactions []model.Transaction) []model.TransactionFormatted {
	formattedTransactions := make([]model.TransactionFormatted, len(allTransactions))

	for i, t := range allTransactions {
		formattedTransactions[i] = FormatTransactionForView(t)
	}

	return formattedTransactions
}

// convert 2022-05-30 string to date obj
func DateStringToDateObj(dateStr string) time.Time {
	dateParts := strings.Split(dateStr, "-")
	datePartsInt := make([]int, len(dateParts))

	for i, value := range dateParts {
		intValue, _ := strconv.Atoi(value)
		datePartsInt[i] = intValue
	}

	year, month, day := datePartsInt[0], time.Month(datePartsInt[1]), datePartsInt[2]

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// Create a model.Summary object for given transactions
func GetSummary(ts []model.Transaction) model.Summary {
	var summary model.Summary

	for _, transaction := range ts {
		if transaction.Type == config.TypeIncomeID {
			summary.Balance += transaction.Amount
			summary.Income += transaction.Amount
		} else {
			summary.Balance -= transaction.Amount
			summary.Expense += transaction.Amount
		}
	}

	return summary
}

func GetNewestAndOldestTDates() (time.Time, time.Time, error) {
	// for oldest
	findOptions := options.Find()
	findOptions.SetSort(bson.M{config.DateID: 1})
	findOptions.SetLimit(1)

	cursor, err := config.Transactions.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		fmt.Println(err.Error())
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())
	}

	oldestT := make([]model.Transaction, len(results))

	for i, resultItem := range results {
		oldestT[i] = BsonDocToTransaction(resultItem)
	}

	// for newest

	findOptions.SetSort(bson.M{config.DateID: -1})
	findOptions.SetLimit(1)

	cursor, err = config.Transactions.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		fmt.Println(err.Error())
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())
	}

	newestT := make([]model.Transaction, len(results))

	for i, resultItem := range results {
		newestT[i] = BsonDocToTransaction(resultItem)
	}

	if len(oldestT) == 0 || len(newestT) == 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("no documents found")
	}

	return oldestT[0].Date, newestT[0].Date, nil
}
