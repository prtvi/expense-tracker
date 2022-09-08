package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	config "github.com/prtvi/expense-tracker/config"
	model "github.com/prtvi/expense-tracker/model"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BsonDocToSettings(doc bson.M) model.Settings {
	var settings model.Settings
	docByte, _ := json.Marshal(doc)
	json.Unmarshal(docByte, &settings)

	return settings
}

func ParseSettingsUrlParams(c echo.Context) model.Settings {
	currency := c.QueryParam(config.CurrencyID)
	monthlyBudget, _ := strconv.ParseFloat(c.QueryParam(config.MonthlyBudgetID), 32)

	modesOfPayment := config.AllModesOfPayment

	for key := range c.QueryParams() {
		if key == config.CurrencyID || key == config.MonthlyBudgetID {
			continue
		}

		modesOfPayment[key] = model.ModeValues{Value: modesOfPayment[key].Value, IsChecked: true}
	}

	return model.Settings{Currency: currency, ModesOfPayment: modesOfPayment, CurrentMonthBudget: float32(monthlyBudget)}
}

func GetCurrencyAndModesOfPayment() (string, map[string]model.ModeValues) {
	cursor := config.Settings.FindOne(context.TODO(), bson.D{})

	fetchedDoc := bson.M{}
	decodeErr := cursor.Decode(&fetchedDoc)
	if decodeErr != nil {
		fmt.Println("No settings document found! Returning empty uninitialized document.")
	}

	settings := BsonDocToSettings(fetchedDoc)

	// duplicating the original const
	mopFiltered := make(map[string]model.ModeValues, len(config.AllModesOfPayment))

	for key, value := range config.AllModesOfPayment {
		mopFiltered[key] = value
	}

	// if the value.IsChecked is true then add to map
	for key, value := range settings.ModesOfPayment {
		modeValues := value

		if value.IsChecked {
			modeValues.IsChecked = true
			mopFiltered[key] = modeValues
		}
	}

	return settings.Currency, mopFiltered
}

func InitAndUpdateSettings(settings model.Settings) error {
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
		return cursor.Err()
	}

	return nil
}
