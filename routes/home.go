package routes

import (
	"net/http"
	"time"

	config "prtvi/expense-tracker/config"
	model "prtvi/expense-tracker/model"
	utils "prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/" route, gets all transactions (formatted for view) and makes a summary to render on page

func Home(c echo.Context) error {
	// get filter params from sort middleware
	sort := c.Get(config.Sort).(string)

	// view params from range of dates
	view := c.Get(config.View).(string)
	viewStartDate := c.Get(config.ViewStartDate).(time.Time)
	viewEndDate := c.Get(config.ViewEndDate).(time.Time)

	var tsForView []model.Transaction
	var tsForViewSummary model.Summary

	ifSubSummary := false

	allTs := utils.GetAllTransactions(sort)
	IfSetup := len(allTs) == 0

	allTSummary := utils.UpdateMainSummary(allTs)

	if view == config.ViewAll {
		tsForView = allTs
	} else {
		tsForView = utils.GetTransactionsByDate(viewStartDate, viewEndDate, sort)
		tsForViewSummary = utils.GetSummary(tsForView)

		if len(tsForView) != 0 {
			ifSubSummary = true
		}

		allTSummary = utils.FetchMainSummary()
	}

	// format the transactions for view (to display on UI)
	formattedTransactions := utils.FormatTransactionsForView(tsForView)

	// budget
	budget := utils.EvalBudget()

	// modes and currency
	// ModesOfPaymentFiltered -> with model.ModeValues{Value, IsChecked: true|false}
	Currency, ModesOfPaymentFiltered := utils.GetCurrencyAndModesOfPayment()

	// fmt.Println(len(ModesOfPayment))

	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"IfSetup": IfSetup,
		// ADD page

		// t-form
		"DateID":             config.DateID,             // 1
		"DescID":             config.DescID,             // 2
		"AmountID":           config.AmountID,           // 3
		"ModeID":             config.ModeID,             // 4
		"TypeInputGroupName": config.TypeInputGroupName, // 5
		"PaidToID":           config.PaidToID,           // 6

		// t-form "mode" (4) input values & text
		"ModesOfPayment": ModesOfPaymentFiltered,

		// t-form "type" (5) ids & values
		"TypeIncomeID":  config.TypeIncomeID,
		"TypeExpenseID": config.TypeExpenseID,

		"Currency": Currency,

		//
		//

		// REPORT page

		// sort-form options
		"ViewID":      config.ViewID,
		"ViewOptions": config.ViewOptions,

		// type select for asc/des sort
		"SortID":    config.Sort,
		"SortAscID": config.SortAscID,
		"SortDesID": config.SortDesID,

		// custom dates container
		"CustomDateStartID": config.CustomDateStartID,
		"CustomDateEndID":   config.CustomDateEndID,

		// main summary
		"TotalIncome":         allTSummary.TotalIncome,
		"TotalExpense":        allTSummary.TotalExpense,
		"TotalBalance":        allTSummary.TotalBalance,
		"SummaryBalanceClass": utils.GetClassNameByValue(allTSummary.TotalBalance),

		// budget summary
		"Budget":               budget.Budget,
		"Spent":                budget.Spent,
		"Remaining":            budget.Remaining,
		"BudgetRemainingClass": utils.GetClassNameByValue(budget.Remaining),

		// if no transactions to show
		"IfNoTransactionsInRange": len(tsForView) == 0,

		// table
		"Transactions": formattedTransactions,

		// sub-summary (for filtered transactions)
		"IfSubSummary":       ifSubSummary,
		"SubIncome":          tsForViewSummary.TotalIncome,
		"SubExpense":         tsForViewSummary.TotalExpense,
		"SubDifference":      tsForViewSummary.TotalBalance,
		"SubDifferenceClass": utils.GetClassNameByValue(tsForViewSummary.TotalBalance),

		// for show range container - filtered, dates
		"ShowingFromDate": utils.FormatDateLong(viewStartDate),
		"ShowingToDate":   utils.FormatDateLong(viewEndDate),

		//
		//

		// SETTINGS page

		"CurrencyID":       config.CurrencyID,
		"MonthlyBudgetID":  config.MonthlyBudgetID,
		"ModesOfPaymentID": config.ModesOfPaymentID,

		"AllModesOfPayment": ModesOfPaymentFiltered,
	})
}
