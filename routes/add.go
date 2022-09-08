package routes

import (
	utils "github.com/prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/add" route, accessible by only javascript
// adds a new transaction to the db

func AddT(c echo.Context) error {
	res := utils.GetResponseMessage(false)

	transaction, err := utils.InitTransaction(c)
	if err != nil {
		return c.JSON(res.StatusCode, res)
	}

	err = utils.InsertTransaction(transaction)
	if err != nil {
		return c.JSON(res.StatusCode, res)
	}

	res = utils.GetResponseMessage(true)
	return c.JSON(res.StatusCode, res)
}
