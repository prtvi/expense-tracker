package utils

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	config "webdev/config"
	model "webdev/model"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// create a new transaction (primitive.D) document from request url data

func InitTransaction(c echo.Context) (bson.D, error) {
	transaction := bson.D{}

	for key, value := range c.QueryParams() {
		// to leave out _id field during update process
		if key == "id" {
			continue
		}

		// to convert amount to number
		if key == config.AmountID {
			valueFloat, err := strconv.ParseFloat(value[0], 32)
			if err != nil {
				return bson.D{}, err
			}

			transaction = append(transaction, bson.E{Key: key, Value: valueFloat})
			continue
		}

		// to enter time.Date object into db
		if key == config.DateID {
			dateParts := strings.Split(value[0], "-")
			datePartsInt := make([]int, len(dateParts))

			for i, value := range dateParts {
				intValue, _ := strconv.Atoi(value)
				datePartsInt[i] = intValue
			}

			year, month, day := datePartsInt[0], time.Month(datePartsInt[1]), datePartsInt[2]

			date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

			transaction = append(transaction, bson.E{Key: key, Value: date})
			continue
		}

		transaction = append(transaction, bson.E{Key: key, Value: value[0]})
	}

	return transaction, nil
}

// returns a model.Transaction obj by id

func GetDocumentById(id string) (model.Transaction, error) {
	// convert string id to object.id
	docID, _ := primitive.ObjectIDFromHex(id)

	cursor := config.Transactions.FindOne(context.TODO(), bson.D{{Key: "_id", Value: docID}})

	fetchedDoc := bson.M{}
	decodeErr := cursor.Decode(&fetchedDoc)
	if decodeErr != nil {
		return model.Transaction{}, decodeErr
	}

	transaction := BsonDocToTransaction(fetchedDoc)
	return transaction, nil
}

// returns an array of model.Transaction from database

func GetTransactions() []model.Transaction {
	cursor, err := config.Transactions.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err.Error())
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())
	}

	allTransactions := make([]model.Transaction, len(results))

	for i, resultItem := range results {
		allTransactions[i] = BsonDocToTransaction(resultItem)
	}

	return allTransactions
}

// insert transaction document to database

func InsertTransaction(transaction bson.D) error {
	_, err := config.Transactions.InsertOne(context.TODO(), transaction)
	if err != nil {
		return err
	}
	return nil
}

// update a transaction, given the string id and toUpdateTransaction

func UpdateTransaction(id string, toUpdateTransaction bson.D) error {
	// convert string id to object.id
	docID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: docID}}
	update := bson.M{"$set": toUpdateTransaction}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	cursor := config.Transactions.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if cursor.Err() != nil {
		return cursor.Err()
	}

	return nil
}

// delete one

func DeleteTransaction(id string) error {
	// convert string id to object.id
	docID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: docID}}
	_, err := config.Transactions.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}
