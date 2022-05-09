package utils

import (
	"encoding/json"
	"time"

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

func FormatDateAndDesc(allTransactions []model.Transaction) []model.Transaction {
	formattedTransactions := make([]model.Transaction, len(allTransactions))

	for i, t := range allTransactions {
		T := t

		// truncating desc text
		if len(T.Desc) > 20 {
			T.Desc = T.Desc[0:20] + "..."
		}

		// formatting date
		timeParsed, _ := time.Parse("2006-01-02", T.Date)
		T.Date = timeParsed.Format(time.RFC1123)[0:16]

		formattedTransactions[i] = T
	}

	return formattedTransactions
}
