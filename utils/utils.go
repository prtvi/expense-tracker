package utils

import (
	"encoding/json"

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

func CreateResponseMessage(statusCode int, success bool) model.ResponseMsg {
	return model.ResponseMsg{StatusCode: statusCode, Success: success}
}
