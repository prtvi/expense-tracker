package utils

import (
	"context"
	"fmt"

	config "webdev/config"

	"go.mongodb.org/mongo-driver/bson"
)

// insert transaction document to database

func InsertTransaction(transaction bson.D) {
	_, err := config.Transactions.InsertOne(context.TODO(), transaction)
	if err != nil {
		fmt.Println(err.Error())
	}
}
