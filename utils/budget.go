package utils

import (
	"context"
	"fmt"
	"time"

	config "prtvi/expense-tracker/config"
	model "prtvi/expense-tracker/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// to set budget on the first go, else to update the budget value
func SetBudget(budgetToBeSet float32) error {
	var budget bson.M

	// find for budget doc in db
	err := config.Budget.FindOne(context.TODO(), bson.M{}).Decode(&budget)

	// if not found then insert new budget doc in db
	if err != nil && err == mongo.ErrNoDocuments {

		now := time.Now()
		currentYear, currentMonth, _ := now.Date()

		newBudget := bson.D{
			{Key: "budget", Value: budgetToBeSet},
			{Key: "month", Value: currentMonth},
			{Key: "year", Value: currentYear},
			{Key: "spent", Value: 0},
			{Key: "remaining", Value: budgetToBeSet},
		}

		_, err2 := config.Budget.InsertOne(context.TODO(), newBudget)
		if err2 != nil {
			return err2
		}

		return nil
	} else {
		// if found then update only the new budgetToBeSet

		// create filter, update and options for querying
		filter := bson.M{}
		update := bson.M{
			"$set": bson.D{
				{Key: "budget", Value: budgetToBeSet},
			},
		}
		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}

		cursor := config.Budget.FindOneAndUpdate(context.TODO(), filter, update, &opt)
		if cursor.Err() != nil {
			return cursor.Err()
		}

		return nil
	}
}

func FetchBudget() model.Budget {
	cursor := config.Budget.FindOne(context.TODO(), bson.M{})

	fetchedDoc := bson.M{}
	decodeErr := cursor.Decode(&fetchedDoc)
	if decodeErr != nil {
		fmt.Println("error")
	}

	return BsonDocToBudget(fetchedDoc)
}

func EvalBudget() model.Budget {
	var budget bson.M

	// find for budget doc in db
	err := config.Budget.FindOne(context.TODO(), bson.M{}).Decode(&budget)
	if err != nil {
		fmt.Println("err")
	}

	// convert to budget obj
	budgetObj := BsonDocToBudget(budget)

	// to init budget struct on every eval, to prevent overwriting over the original doc
	newBudget := model.Budget{
		Budget:    budgetObj.Budget,
		Month:     budgetObj.Month,
		Year:      budgetObj.Year,
		Spent:     0,
		Remaining: 0,
	}

	// get the month begin & end date for quering the transactions between that month
	monthBeginDate, monthEndDate := FirstAndLastDayOfMonth(budgetObj.Year, budgetObj.Month, time.Local)
	monthEndDate = LastSecondOfTheDay(monthEndDate)

	// get transactions for that month
	transactionsThisMonth := GetTransactionsByDate(monthBeginDate, monthEndDate, config.SortAscID)

	// update newBudget
	for _, transaction := range transactionsThisMonth {
		if transaction.Type == config.TypeExpenseID {
			newBudget.Spent += transaction.Amount
		}
	}

	newBudget.Remaining = newBudget.Budget - newBudget.Spent

	// update the budget into db

	// create filter, update and options for querying
	filter := bson.M{}
	update := bson.M{
		"$set": bson.D{
			{Key: "spent", Value: newBudget.Spent},
			{Key: "remaining", Value: newBudget.Remaining},
		},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	cursor := config.Budget.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if cursor.Err() != nil {
		fmt.Println("err")
	}

	return newBudget
}
