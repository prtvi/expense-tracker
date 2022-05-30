package routes

import (
	"net/http"
	"time"

	config "webdev/config"
	model "webdev/model"
	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

// "/" route, gets all transactions (formatted for view) and makes a summary to render on page

func Home(c echo.Context) error {
	sortStartDate, sortEndDate := c.Get(config.SortStartDate).(time.Time), c.Get(config.SortEndDate).(time.Time)

	var allTs []model.Transaction
	var allTSummary model.Summary

	var tsForView []model.Transaction
	var tsForViewSummary model.Summary

	if sortStartDate.Year() == 0001 {
		// fetch all transactions & update the summary of transactions (also in db)
		allTs = utils.GetTransactions()
		allTSummary = utils.UpdateSummary(allTs)

		tsForView = allTs
	} else {
		tsForView = utils.GetTransactionsByDate(sortStartDate, sortEndDate)
		tsForViewSummary = utils.GetSummary(tsForView)

		allTSummary = utils.FetchSummary()
	}

	// format the transactions for view (to display on UI)
	formattedTransactions := utils.FormatTransactionsForView(tsForView)

	return c.Render(http.StatusOK, "index", map[string]interface{}{
		// overall summary options
		"TotalIncome":         allTSummary.TotalIncome,
		"TotalExpense":        allTSummary.TotalExpense,
		"CurrentBalance":      allTSummary.CurrentBalance,
		"CurrentBalanceClass": utils.GetClassNameByValue(allTSummary.CurrentBalance),

		// for sorted, dates
		"ShowingFromDate": utils.FormatDateLong(sortStartDate),
		"ShowingToDate":   utils.FormatDateLong(sortEndDate),

		// sub-summary (for sorted transaction options)
		"IfSubSummary":       len(allTs) != len(tsForView),
		"SubTotalIncome":     tsForViewSummary.TotalIncome,
		"SubTotalExpense":    tsForViewSummary.TotalExpense,
		"SubDifference":      tsForViewSummary.CurrentBalance,
		"SubDifferenceClass": utils.GetClassNameByValue(tsForViewSummary.CurrentBalance),

		"IfNoTransactions": len(formattedTransactions) == 0,
		"Transactions":     formattedTransactions,
		"Currency":         "₹",

		// element ids
		// t-form
		"DateID":             config.DateID,             // 1
		"DescID":             config.DescID,             // 2
		"AmountID":           config.AmountID,           // 3
		"ModeID":             config.ModeID,             // 4
		"TypeInputGroupName": config.TypeInputGroupName, // 5
		"PaidToID":           config.PaidToID,           // 6

		// t-form "mode" (4) input values & text
		"Modes": config.Modes,

		// t-form "type" (5) ids & values
		"TypeIncomeID":  config.TypeIncomeID,
		"TypeExpenseID": config.TypeExpenseID,

		// sort-form options
		"SortOptions":       config.SortOptions,
		"SortInputID":       config.SortInputID,
		"CustomDateStartID": config.CustomDateStartID,
		"CustomDateEndID":   config.CustomDateEndID,
	})
}
