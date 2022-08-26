package utils

import (
	"context"
	"encoding/json"
	"fmt"
	config "prtvi/expense-tracker/config"
	model "prtvi/expense-tracker/model"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func BsonDocToSettings(doc bson.M) model.Settings {
	var settings model.Settings
	docByte, _ := json.Marshal(doc)
	json.Unmarshal(docByte, &settings)

	return settings
}

func ParseSettingsUrlParams(c echo.Context) model.Settings {

	currency := c.QueryParam(config.CurrencyID)
	modesOfPayment := strings.Split(c.QueryParam(config.ModesOfPaymentID), ",")
	for i := range modesOfPayment {
		modesOfPayment[i] = strings.ToLower(strings.TrimSpace(modesOfPayment[i]))
	}
	monthlyBudget, _ := strconv.ParseFloat(c.QueryParam(config.MonthlyBudgetID), 32)

	return model.Settings{Currency: currency, ModesOfPayment: modesOfPayment, CurrentMonthBudget: float32(monthlyBudget)}
}

func GetCurrencyAndModesOfPayment() (string, []string) {
	cursor := config.Settings.FindOne(context.TODO(), bson.D{})

	fetchedDoc := bson.M{}
	decodeErr := cursor.Decode(&fetchedDoc)
	if decodeErr != nil {
		fmt.Println("decode error")
	}

	settings := BsonDocToSettings(fetchedDoc)

	return settings.Currency, settings.ModesOfPayment
}

func UpdateSettings(settings model.Settings) bool {
	// create filter, update and options for querying
	filter := bson.M{}
	update := bson.M{
		"$set": bson.D{
			{Key: "currency", Value: settings.Currency},
			{Key: "modes_of_payment", Value: settings.ModesOfPayment},
			{Key: "current_month_budget", Value: settings.CurrentMonthBudget},
		},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	cursor := config.Settings.FindOneAndUpdate(context.TODO(), filter, update, &opt)

	if cursor.Err() != nil {
		return false
	}

	return true
}
