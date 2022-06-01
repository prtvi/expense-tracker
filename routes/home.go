package routes

import (
	"net/http"
	"time"

	config "webdev/config"
	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

// "/" route, gets all transactions (formatted for view) and makes a summary to render on page

func Home(c echo.Context) error {
	sortStartDate := c.Get(config.SortStartDate).(time.Time)
	sortEndDate := c.Get(config.SortEndDate).(time.Time)

	sortBy := c.Get(config.SortBy).(string)
	sortFor := c.Get(config.SortFor).(string)

	// init all arrays

	allTs := utils.GetAllTransactions(sortBy)
	allTSummary := utils.UpdateSummary(allTs)

	tsForView := utils.GetTransactionsByDate(sortStartDate, sortEndDate, sortBy)
	tsForViewSummary := utils.GetSummary(tsForView)

	if sortFor == config.SortAllValue {
		tsForView = allTs
	} else {
		allTSummary = utils.FetchSummary()
	}

	// format the transactions for view (to display on UI)
	formattedTransactions := utils.FormatTransactionsForView(tsForView)

	return c.Render(http.StatusOK, "index", map[string]interface{}{
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

		"Currency": "₹",

		// is true if there are 0 transactions in the entire db
		"IfZeroTransactions": len(allTs) == 0 && len(tsForView) == 0,

		// main summary
		"TotalIncome":         allTSummary.TotalIncome,
		"TotalExpense":        allTSummary.TotalExpense,
		"CurrentBalance":      allTSummary.CurrentBalance,
		"CurrentBalanceClass": utils.GetClassNameByValue(allTSummary.CurrentBalance),

		// transactions to show
		"IfNoTransactionToView": len(formattedTransactions) == 0,
		"Transactions":          formattedTransactions,

		// for sorted, dates
		"ShowingFromDate": utils.FormatDateLong(sortStartDate),
		"ShowingToDate":   utils.FormatDateLong(sortEndDate),

		// sub-summary (for sorted transactions)
		"IfSubSummary":       len(allTs) != len(tsForView) && len(formattedTransactions) != 0,
		"SubTotalIncome":     tsForViewSummary.TotalIncome,
		"SubTotalExpense":    tsForViewSummary.TotalExpense,
		"SubDifference":      tsForViewSummary.CurrentBalance,
		"SubDifferenceClass": utils.GetClassNameByValue(tsForViewSummary.CurrentBalance),

		// sort-form options
		"SortForID":         config.SortForID,
		"SortOptions":       config.SortOptions,
		"CustomDateStartID": config.CustomDateStartID,
		"CustomDateEndID":   config.CustomDateEndID,

		// type select for asc/des sort
		"SortByID":    config.SortByID,
		"SortByAscID": config.SortByAscID,
		"SortByDesID": config.SortByDesID,
	})
}
