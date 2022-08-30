package config

import (
	"prtvi/expense-tracker/model"

	"go.mongodb.org/mongo-driver/mongo"
)

// Global mongo collection connection
var Transactions mongo.Collection
var Summary mongo.Collection
var Budget mongo.Collection
var Settings mongo.Collection

// for date formatting
const MAX_DESC_LEN = 20

// for format: 2022-05-25 (short date)
const FORMAT_DATE_STR_LEN = 10

// for format: Wed, 25 May (in words)
const FORMAT_DATE_STR_LEN_WORDS = 11

// for format: Wed, 25 May 2022 (long date)
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

// t-form "type" (5) ids & values
// if changed, then gotta change the value for line with content "{{ if eq .Type "income" }}" in index.html
const TypeIncomeID = "income"
const TypeExpenseID = "expense"

//
//
//
//

// sort-form
// sort-form ids
const ViewID = "view"

// sort-form select option element values (view)
const ViewAll = "1_all"
const ViewLast7Days = "2_last7"
const ViewLast30Days = "3_last30"
const ViewThisMonth = "4_this_month"
const ViewLastMonth = "5_last_month"

// for when view element is chosen as "custom"
const ViewCustom = "6_custom"

const CustomDateStartID = "date_start"
const CustomDateEndID = "date_end"

// type select
const SortID = "sort"

// id and values
const SortAscID = "asc"
const SortDesID = "des"

// keys for passing data between middlewares
const Sort = SortID

const View = ViewID
const ViewStartDate = "view_start_date"
const ViewEndDate = "view_end_date"

// numbers given to sort them according to preference
var ViewOptions map[string]string = map[string]string{
	// select option element value : text shown on elements on UI
	ViewAll:        "All",
	ViewLast7Days:  "Last 7 days",
	ViewLast30Days: "Last 30 days",
	ViewThisMonth:  "This month",
	ViewLastMonth:  "Last month",
	ViewCustom:     "Custom",
}

// settings page

const CurrencyID = "currency"
const MonthlyBudgetID = "monthly_budget"
const ModesOfPaymentID = "modes_of_payment"

const Cash = "1_cash"
const Phonepe = "2_phonepe"
const Googlepay = "3_googlepay"
const Paytm = "4_paytm"
const Card = "5_card"
const Other = "6_other"

var AllModesOfPayment map[string]model.ModeValues = map[string]model.ModeValues{
	// select option element value : text shown on elements on UI
	Cash:      {Value: "Cash", IsChecked: false},
	Phonepe:   {Value: "PhonePe", IsChecked: false},
	Googlepay: {Value: "Google Pay", IsChecked: false},
	Paytm:     {Value: "PayTM", IsChecked: false},
	Card:      {Value: "Card", IsChecked: false},
	Other:     {Value: "Other", IsChecked: false},
}
