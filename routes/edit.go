package routes

import (
	utils "github.com/prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/edit" route, accessible by only javascript
// updates the document with input form data

func EditT(c echo.Context) error {
	id := c.QueryParam("id")
	res := utils.GetResponseMessage(false)

	updatedTransaction, err := utils.InitTransaction(c)
	if err != nil {
		return c.JSON(res.StatusCode, res)
	}

	err = utils.UpdateTransaction(id, updatedTransaction)
	if err != nil {
		return c.JSON(res.StatusCode, res)
	}

	res = utils.GetResponseMessage(true)
	return c.JSON(res.StatusCode, res)
}
