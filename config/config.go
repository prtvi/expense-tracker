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

// element ids
// t-form
const DateID = "date"             // 1. t-date                 (type date)
const DescID = "desc"             // 2. t-description          (type text)
const AmountID = "amount"         // 3. t-amount               (type number)
const ModeID = "mode"             // 4. mode-of-t              (type select dropdown)
const TypeInputGroupName = "type" // 5. t-type, income|expense (type radio group)
const PaidToID = "paid_to"        // 6. amount paid to         (type text)

// t-form "mode" (4) input values & text
var Modes = []string{"PhonePe", "Google Pay", "Cash", "PayTM"}

// t-form "type" (5) ids & values
// if changed, then gotta change the value for line with content "{{ if eq .Type "income" }}" in index.html
const TypeIncomeID = "income"
const TypeExpenseID = "expense"

// sort-form ids
const SortInputID = "sort"
const CustomDateStartID = "date_start"
const CustomDateEndID = "date_end"

// sort-form select option element values
const SortAllValue = "1_all"
const SortLast7DaysValue = "2_last_seven"
const SortLast30DaysValue = "3_last_thirty"
const SortThisMonthValue = "4_this_month"
const SortLastMonthValue = "5_last_month"
const SortCustomValue = "6_custom"

// numbers given to sort them according to preference
var SortOptions map[string]string = map[string]string{
	// select option element value : text shown on elements on UI
	SortAllValue:        "All",
	SortLast7DaysValue:  "Last 7 days",
	SortLast30DaysValue: "Last 30 days",
	SortThisMonthValue:  "This month",
	SortLastMonthValue:  "Last month",
	SortCustomValue:     "Custom",
}
