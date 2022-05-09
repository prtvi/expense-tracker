package routes

import (
	"fmt"
	"net/http"

	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

// "/" route, gets all transactions and makes a summary to render on page

func Home(c echo.Context) error {
	fmt.Println("hit: GET: /")

	allTransactions := utils.GetTransactions()
	summary := utils.UpdateSummary(allTransactions)
	formattedTransactions := utils.FormatDateAndDesc(allTransactions)

	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"totalExpense":     summary.TotalExpense,
		"totalIncome":      summary.TotalIncome,
		"currentBalance":   summary.CurrentBalance,
		"transactions":     formattedTransactions,
		"ifNoTransactions": len(allTransactions) == 0,
	})
}
