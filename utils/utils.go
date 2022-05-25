package utils

import (
	"encoding/json"
	"time"

	model "webdev/model"

	"go.mongodb.org/mongo-driver/bson"
)

const MAX_DESC_LEN = 20

// for format: 2022-05-25
const FORMAT_DATE_STR_LEN = 10

// for format: Wed, 25 May 2022
const FORMAT_DATE_STR_LEN_LONG = 16

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

// convert model.Transaction to model.TransactionFormatted

func FormatTransaction(t model.Transaction) model.TransactionFormatted {
	// date object to format: 2022-05-25
	return model.TransactionFormatted{
		ID:     t.ID,
		Date:   t.Date.String()[0:FORMAT_DATE_STR_LEN],
		Desc:   t.Desc,
		Amount: t.Amount,
		Type:   t.Type,
		Mode:   t.Mode,
		PaidTo: t.PaidTo,
	}
}

// specifically for "/get" route, to format model.Transaction to model.TransactionFormatted, truncate desc text to MAX_DESC_LEN & format date to format: Wed, 25 May 2022

func FormatDateAndDesc(allTransactions []model.Transaction) []model.TransactionFormatted {
	formattedTransactions := make([]model.TransactionFormatted, len(allTransactions))

	for i, t := range allTransactions {
		T := FormatTransaction(t)

		// truncating desc text (only when viewed inside table)
		if len(t.Desc) > MAX_DESC_LEN {
			T.Desc = t.Desc[0:MAX_DESC_LEN] + "..."
		}

		// to format into: Wed, 25 May 2022
		T.Date = t.Date.Format(time.RFC1123Z)[0:FORMAT_DATE_STR_LEN_LONG]

		formattedTransactions[i] = T
	}

	return formattedTransactions
}
