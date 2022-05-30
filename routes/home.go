package routes

import (
	"fmt"
	"net/http"
	"time"

	"webdev/config"
	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

// "/" route, gets all transactions (formatted for view) and makes a summary to render on page

func Home(c echo.Context) error {
	sortStartDate, sortEndDate := c.Get(config.SortStartDate).(time.Time), c.Get(config.SortEndDate).(time.Time)

	sortedTransactions := utils.GetTransactionsByDate(sortStartDate, sortEndDate)

	fmt.Println(sortedTransactions)

	// fetch all transactions
	allTransactions := utils.GetTransactions()

	// update the summary of transactions (also in db)
	summary := utils.UpdateSummary(allTransactions)

	// format the transactions for view (to display on UI)
	formattedTransactions := utils.FormatTransactionsForView(allTransactions)

	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"TotalIncome":  summary.TotalIncome,
		"TotalExpense": summary.TotalExpense,

		"CurrentBalance":      summary.CurrentBalance,
		"CurrentBalanceClass": utils.GetClassNameByValue(summary.CurrentBalance),

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
