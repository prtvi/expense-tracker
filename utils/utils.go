package utils

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	config "webdev/config"
	model "webdev/model"

	"go.mongodb.org/mongo-driver/bson"
)

// convert primitive.M to model.Transaction object

func BsonDocToTransaction(doc bson.M) model.Transaction {
	var transaction model.Transaction
	docByte, _ := json.Marshal(doc)
	json.Unmarshal(docByte, &transaction)

	return transaction
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

// converts a single transaction from model.Transaction to model.TransactionFormatted (not for view, to be loaded on t-form)

func FormatTransaction(t model.Transaction) model.TransactionFormatted {
	// date object to format: 2022-05-25
	return model.TransactionFormatted{
		ID:     t.ID,
		Date:   t.Date.String()[0:config.FORMAT_DATE_STR_LEN],
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
	T.Date = t.Date.Format(time.RFC1123Z)[0:config.FORMAT_DATE_STR_LEN_LONG]

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
