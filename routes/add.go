package routes

import (
	"net/http"

	utils "prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/add" route, accessible by only javascript
// adds a new transaction to the db

func AddT(c echo.Context) error {
	transaction, err := utils.InitTransaction(c)

	res := utils.CreateResponseMessage(http.StatusBadRequest, false, "Operation failed")

	if err != nil {
		return c.JSON(http.StatusBadRequest, res)
	}

	err = utils.InsertTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, res)
	}

	res = utils.CreateResponseMessage(http.StatusOK, true, "Success")
	return c.JSON(http.StatusOK, res)
}
