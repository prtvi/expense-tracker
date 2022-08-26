package utils

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	config "prtvi/expense-tracker/config"
	model "prtvi/expense-tracker/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// constructor for generating model.ResponseMessage objects

func CreateResponseMessage(statusCode int, success bool, message string) model.ResponseMsg {
	return model.ResponseMsg{StatusCode: statusCode, Success: success, Message: message}
}

func GetClassNameByValue(value float32) string {
	if value >= 0 {
		return config.ClassTTypeIncome
	}
	return config.ClassTTypeExpense
}

// date object to date string with format: 2022-05-25
func FormatDateShort(d time.Time) string {
	return d.String()[0:config.FORMAT_DATE_STR_LEN]
}

// date object to date string with format: Wed, 25 May
func FormatDateWords(d time.Time) string {
	return d.Format(time.RFC1123Z)[0:config.FORMAT_DATE_STR_LEN_WORDS]
}

// date object to date string with format: Wed, 25 May 2022
func FormatDateLong(d time.Time) string {
	return d.Format(time.RFC1123Z)[0:config.FORMAT_DATE_STR_LEN_LONG]
}

// convert 2022-05-30 string to date obj
func DateStringToDateObj(dateStr string, insert bool) time.Time {
	dateParts := strings.Split(dateStr, "-")
	datePartsInt := make([]int, len(dateParts))

	for i, value := range dateParts {
		intValue, _ := strconv.Atoi(value)
		datePartsInt[i] = intValue
	}

	year, month, day := datePartsInt[0], time.Month(datePartsInt[1]), datePartsInt[2]

	if insert {
		// if the converted time is to be inserted into the db, then add time to the date as well

		now := time.Now()
		currentLocation := now.Location()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		return time.Date(year, month, day, hour, min, sec, 0, currentLocation)
	}

	// else do not add time to date
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func GetNewestAndOldestTDates() (time.Time, time.Time, error) {
	// for oldest
	findOptions := options.Find()
	findOptions.SetSort(bson.M{config.DateID: 1})
	findOptions.SetLimit(1)

	cursor, err := config.Transactions.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		fmt.Println(err.Error())
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())
	}

	oldestT := make([]model.Transaction, len(results))

	for i, resultItem := range results {
		oldestT[i] = BsonDocToTransaction(resultItem)
	}

	// for newest

	findOptions.SetSort(bson.M{config.DateID: -1})
	findOptions.SetLimit(1)

	cursor, err = config.Transactions.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		fmt.Println(err.Error())
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())
	}

	newestT := make([]model.Transaction, len(results))

	for i, resultItem := range results {
		newestT[i] = BsonDocToTransaction(resultItem)
	}

	if len(oldestT) == 0 || len(newestT) == 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("no documents found")
	}

	return oldestT[0].Date, newestT[0].Date, nil
}

// get the first & last day of the month, without time (time->0)
func FirstAndLastDayOfMonth(year, month int, loc *time.Location) (time.Time, time.Time) {
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return firstOfMonth, lastOfMonth
}

// convert a time from 00:00:00 to 23:59:59.9999
func LastSecondOfTheDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1).Add(-time.Nanosecond)
}
