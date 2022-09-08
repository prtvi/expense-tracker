package routes

import (
	"net/http"

	utils "github.com/prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/get" route, accessible by only javascript
// returns a single formatted transaction by doc_id (formatted not for view)

func ReturnT(c echo.Context) error {
	id := c.QueryParam("id")

	transaction, err := utils.GetDocumentById(id)
	if err != nil {
		res := utils.GetResponseMessage(false)
		return c.JSON(res.StatusCode, res)
	}

	// format transaction to be loaded on t-form (not for view)
	transactionFormatted := utils.FormatTransaction(transaction)

	return c.JSON(http.StatusOK, transactionFormatted)
}
