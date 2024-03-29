package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	config "github.com/prtvi/expense-tracker/config"
	model "github.com/prtvi/expense-tracker/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BsonDocToBudget(doc bson.M) model.Budget {
	var budget model.Budget
	docByte, _ := json.Marshal(doc)
	json.Unmarshal(docByte, &budget)

	return budget
}

// sets all values whenever called
func InitAndSetBudget(budgetToBeSet float32) error {
	// if found then update only the new budgetToBeSet
	currentMonth, currentYear := GetCurrentMonthAndYear()

	// create filter, update and options for querying
	filter := bson.M{}
	update := bson.M{
		"$set": bson.D{
			{Key: "budget", Value: budgetToBeSet},
			{Key: "month", Value: currentMonth},
			{Key: "year", Value: currentYear},
			{Key: "spent", Value: 0},
			{Key: "remaining", Value: budgetToBeSet},
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
		fmt.Println("No budget document found!")
		return model.Budget{}
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
	monthEndDate = GoToLastSecondOfTheDay(monthEndDate)

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
		fmt.Println("No budget document found to update")
		return model.Budget{}
	}

	return newBudget
}
