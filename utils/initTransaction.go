package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// create & insert a new transaction document from request url data

func InitTransaction(c echo.Context) (bson.D, error) {
	transaction := bson.D{}

	// fmt.Println(c.Request().URL)

	for key, value := range c.QueryParams() {
		if key != "amount" {
			transaction = append(transaction, bson.E{Key: key, Value: value[0]})
			continue
		}

		valueFloat, err := strconv.ParseFloat(value[0], 32)
		if err != nil {
			return bson.D{}, err
		}
		transaction = append(transaction, bson.E{Key: key, Value: valueFloat})
	}

	return transaction, nil
}
