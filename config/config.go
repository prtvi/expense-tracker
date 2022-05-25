package config

import "go.mongodb.org/mongo-driver/mongo"

// Global mongo collection connection
var Transactions mongo.Collection
var Summary mongo.Collection

// for date formatting
const MAX_DESC_LEN = 20

// for format: 2022-05-25
const FORMAT_DATE_STR_LEN = 10

// for format: Wed, 25 May 2022
const FORMAT_DATE_STR_LEN_LONG = 16

// class names
const ClassTTypeIncome = "t-type-income"
const ClassTTypeExpense = "t-type-expense"
