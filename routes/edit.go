package routes

import (
	"net/http"

	utils "prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/edit" route, accessible by only javascript
// updates the document with input form data

func EditT(c echo.Context) error {
	updatedTransaction, err := utils.InitTransaction(c)
	if err != nil {
		res := utils.CreateResponseMessage(http.StatusBadRequest, false, "Operation failed")
		return c.JSON(http.StatusBadRequest, res)
	}

	id := c.QueryParam("id")

	err = utils.UpdateTransaction(id, updatedTransaction)
	if err != nil {
		res := utils.CreateResponseMessage(http.StatusBadRequest, false, "Operation failed")
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.CreateResponseMessage(http.StatusOK, true, "Success")
	return c.JSON(http.StatusOK, res)
}
