package routes

import (
	"fmt"
	"net/http"

	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

// "/edit" route, accessible by only javascript
// updates the document with input form data

func EditT(c echo.Context) error {
	fmt.Println("hit: GET: /edit")

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
