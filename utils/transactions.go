package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	config "github.com/prtvi/expense-tracker/config"
	model "github.com/prtvi/expense-tracker/model"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// convert primitive.M to model.Transaction object

func BsonDocToTransaction(doc bson.M) model.Transaction {
	var transaction model.Transaction
	docByte, _ := json.Marshal(doc)
	json.Unmarshal(docByte, &transaction)

	return transaction
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

// formats a single transaction for view, model.Transaction to model.TransactionFormatted (format for view)
// truncate desc text to MAX_DESC_LEN & format date to format: Wed, 25 May
// to view only on table
func FormatTransactionForView(t model.Transaction) model.TransactionFormatted {
	T := FormatTransaction(t)

	// 1. format date to format into: Wed, 25 May
	T.Date = FormatDateWords(t.Date)

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

		// TODO: for update process, change date only, not time of insertion
		// to enter time.Date object into db
		if key == config.DateID {
			date := DateStringToDatetimeObj(value[0])

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

// get sort options obj based on asc/des sort
func getSortOptions(sort string) *options.FindOptions {
	findOptions := options.Find()

	if sort == config.SortDesID {
		findOptions.SetSort(bson.M{config.DateID: -1}) // descending
	} else {
		findOptions.SetSort(bson.M{config.DateID: 1}) // ascending (by default)
	}

	return findOptions
}

// returns an array of model.Transaction from database

func GetAllTransactions(sort string) []model.Transaction {
	findOptions := getSortOptions(sort)

	cursor, err := config.Transactions.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		fmt.Println("No transactions found!", err.Error())
		return []model.Transaction{}
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("No transactions found!", err.Error())
		return []model.Transaction{}
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

// get transactions within date range

func GetTransactionsByDate(startDate, endDate time.Time, sort string) []model.Transaction {
	filter := bson.M{config.DateID: bson.M{"$gte": startDate, "$lte": endDate}}
	findOptions := getSortOptions(sort)

	cursor, err := config.Transactions.Find(context.TODO(), filter, findOptions)
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

// count the number of transactions in the db
func CountTransactions() int64 {
	count, _ := config.Transactions.CountDocuments(context.TODO(), bson.D{})
	return count
}
