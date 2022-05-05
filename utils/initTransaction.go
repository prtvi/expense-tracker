package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// create & insert a new transaction document from form data

func InitTransaction(c echo.Context) bson.D {
	transaction := bson.D{}
	formValues, _ := c.FormParams()

	for key, value := range formValues {
		if key != "amount" {
			transaction = append(transaction, bson.E{Key: key, Value: value[0]})
			continue
		}
		valueFloat, _ := strconv.ParseFloat(value[0], 32)
		transaction = append(transaction, bson.E{Key: key, Value: valueFloat})
	}

	return transaction
}
