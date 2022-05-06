package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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
		if key == "amount" {
			valueFloat, err := strconv.ParseFloat(value[0], 32)
			if err != nil {
				return bson.D{}, err
			}

			transaction = append(transaction, bson.E{Key: key, Value: valueFloat})
			continue
		}

		transaction = append(transaction, bson.E{Key: key, Value: value[0]})
	}

	return transaction, nil
}
