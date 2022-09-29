package routes

import (
	"net/http"
	"time"

	config "github.com/prtvi/expense-tracker/config"
	model "github.com/prtvi/expense-tracker/model"
	utils "github.com/prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/" route, gets all transactions (formatted for view) and makes a summary to render on page

func Home(c echo.Context) error {
	// GET VALUES FROM MIDDLEWARE
	sort := c.Get(config.Sort).(string)
	view := c.Get(config.View).(string)
	viewStartDate := c.Get(config.ViewStartDate).(time.Time)
	viewEndDate := c.Get(config.ViewEndDate).(time.Time)

	// INIT global varibles

	var tsForView []model.Transaction
	var tsForViewSummary model.Summary

	allTs := utils.GetAllTransactions(sort)
	var allTSummary model.Summary
	var FormattedTs []model.TransactionFormatted

	var IfTransactions bool = len(allTs) > 0
	var IfSubSummary bool = false

	Budget := utils.EvalBudget()
	Month, Year := utils.GetCurrentMonthAndYear()
	Currency, MopFiltered := utils.GetCurrencyAndModesOfPayment()

	OldestT := utils.GetDateObj(Year)
	NewestT := utils.GetDateObj(Year)

	if IfTransactions {
		allTSummary, OldestT, NewestT = utils.UpdateMainSummary()

		if view == config.ViewAll {
			tsForView = allTs
		} else {
			tsForView = utils.GetTransactionsByDate(viewStartDate, viewEndDate, sort)
			tsForViewSummary = utils.GetSummary(viewStartDate, viewEndDate, sort)

			allTSummary = utils.FetchMainSummary()

			if len(tsForView) != 0 {
				IfSubSummary = true
			}
		}

		FormattedTs = utils.FormatTransactionsForView(tsForView)
	}

	return c.Render(http.StatusOK, "index", map[string]interface{}{
		// ------------- CONSTANTS -------------

		// ADD page

		// t-form
		"DateID":             config.DateID,             // 1
		"DescID":             config.DescID,             // 2
		"AmountID":           config.AmountID,           // 3
		"ModeID":             config.ModeID,             // 4
		"TypeInputGroupName": config.TypeInputGroupName, // 5
		"PaidToID":           config.PaidToID,           // 6

		// t-form "type" (5) ids & values
		"TypeIncomeID":  config.TypeIncomeID,
		"TypeExpenseID": config.TypeExpenseID,

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

		//
		//

		// SETTINGS page

		"CurrencyID":       config.CurrencyID,
		"MonthlyBudgetID":  config.MonthlyBudgetID,
		"ModesOfPaymentID": config.ModesOfPaymentID,

		//
		//
		//
		//
		//

		// ------------- VARIABLES -------------

		"IfTransactions": IfTransactions,

		// ADD page

		// t-form "mode" (4) input values & text
		"ModesOfPayment": MopFiltered,
		"Currency":       Currency,

		//
		//

		// REPORT page

		// main summary
		"TotalIncome":         allTSummary.TotalIncome,
		"TotalExpense":        allTSummary.TotalExpense,
		"TotalBalance":        allTSummary.TotalBalance,
		"SummaryBalanceClass": utils.GetClassNameByValue(allTSummary.TotalBalance),

		// main summary data-date attributes
		"DateStartSummary": utils.FormatDateShort(OldestT),
		"DateEndSummary":   utils.FormatDateShort(NewestT),

		// budget summary
		"Budget":               Budget.Budget,
		"Spent":                Budget.Spent,
		"Remaining":            Budget.Remaining,
		"BudgetRemainingClass": utils.GetClassNameByValue(Budget.Remaining),

		// if no transactions to show
		"IfNoTransactionsInRange": len(tsForView) == 0,

		// table
		"Transactions": FormattedTs,

		// sub-summary (for filtered transactions)
		"IfSubSummary":       IfSubSummary,
		"SubIncome":          tsForViewSummary.TotalIncome,
		"SubExpense":         tsForViewSummary.TotalExpense,
		"SubDifference":      tsForViewSummary.TotalBalance,
		"SubDifferenceClass": utils.GetClassNameByValue(tsForViewSummary.TotalBalance),

		// sub summary data-date attributes
		"DateStartSubSummary": utils.FormatDateShort(viewStartDate),
		"DateEndSubSummary":   utils.FormatDateShort(viewEndDate),

		// for show range container - filtered, dates
		"ShowingFromDate": utils.FormatDateLong(viewStartDate),
		"ShowingToDate":   utils.FormatDateLong(viewEndDate),

		//
		//

		// SETTINGS page

		"CurrentMonth":      time.Month(Month),
		"AllModesOfPayment": MopFiltered,
	})
}
