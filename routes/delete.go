package routes

import (
	utils "github.com/prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/edit" route, accessible by only javascript
// updates the document with input form data

func DeleteT(c echo.Context) error {
	id := c.QueryParam("id")

	err := utils.DeleteTransaction(id)
	if err != nil {
		res := utils.GetResponseMessage(false)
		return c.JSON(res.StatusCode, res)
	}

	res := utils.GetResponseMessage(true)
	return c.JSON(res.StatusCode, res)
}
