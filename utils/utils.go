package utils

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	config "github.com/prtvi/expense-tracker/config"
	model "github.com/prtvi/expense-tracker/model"
	utilModel "github.com/prtvi/expense-tracker/model/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createResponseMessage(statusCode int, success bool, message string) utilModel.ResponseMsg {
	// constructor for generating utilModel.ResponseMessage objects
	return utilModel.ResponseMsg{StatusCode: statusCode, Success: success, Message: message}
}

func GetClassNameByValue(value float32) string {
	if value >= 0 {
		return config.ClassTTypeIncome
	}
	return config.ClassTTypeExpense
}

func FormatDateShort(d time.Time) string {
	// date object to date string with format: 2022-05-25
	return d.String()[0:config.FORMAT_DATE_STR_LEN]
}

func FormatDateWords(d time.Time) string {
	// date object to date string with format: Wed, 25 May
	return d.Format(time.RFC1123Z)[0:config.FORMAT_DATE_STR_LEN_WORDS]
}

func FormatDateLong(d time.Time) string {
	// date object to date string with format: Wed, 25 May 2022
	return d.Format(time.RFC1123Z)[0:config.FORMAT_DATE_STR_LEN_LONG]
}

func parseDateStr(s string) (int, time.Month, int) {
	dateParts := strings.Split(s, "-")
	datePartsInt := make([]int, len(dateParts))

	for i, value := range dateParts {
		intValue, _ := strconv.Atoi(value)
		datePartsInt[i] = intValue
	}

	return datePartsInt[0], time.Month(datePartsInt[1]), datePartsInt[2]
}

func DateStringToDateObj(dateStr string) time.Time {
	// convert 2022-05-30 string to date obj
	year, month, day := parseDateStr(dateStr)

	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func DateStringToDatetimeObj(dateStr string) time.Time {
	// convert 2022-05-30 string to date obj with time as well
	year, month, day := parseDateStr(dateStr)

	now := time.Now()
	currentLocation := now.Location()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	return time.Date(year, month, day, hour, min, sec, 0, currentLocation)
}

func GetNewestAndOldestTDates(year int) (time.Time, time.Time, error) {
	// for oldest
	findOptions := options.Find()
	findOptions.SetSort(bson.M{config.DateID: 1})
	findOptions.SetLimit(1)

	cursor, err := config.Transactions.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		fmt.Println("No oldest transaction found!", err.Error())
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("No oldest transaction found!", err.Error())
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
		fmt.Println("No newest transaction found!", err.Error())
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("No newest transaction found!", err.Error())
	}

	newestT := make([]model.Transaction, len(results))

	for i, resultItem := range results {
		newestT[i] = BsonDocToTransaction(resultItem)
	}

	if len(oldestT) == 0 || len(newestT) == 0 {
		return GetDateObj(year), GetDateObj(year), fmt.Errorf("No oldest or newest documents found!")
	}

	return oldestT[0].Date, newestT[0].Date, nil
}

func FirstAndLastDayOfMonth(year, month int, loc *time.Location) (time.Time, time.Time) {
	// get the first & last day of the month, without time (time->0)
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return firstOfMonth, lastOfMonth
}

func GoToLastSecondOfTheDay(t time.Time) time.Time {
	// convert a time from 00:00:00 to 23:59:59.9999
	return t.AddDate(0, 0, 1).Add(-time.Nanosecond)
}

func GetCurrentMonthAndYear() (int, int) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()

	return int(currentMonth), currentYear
}

func GetDateObj(year int) time.Time {
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
}

func GetResponseMessage(success bool) utilModel.ResponseMsg {
	if success {
		return createResponseMessage(http.StatusOK, true, "Success")
	} else {
		return createResponseMessage(http.StatusBadRequest, false, "Operation failed")
	}
}
