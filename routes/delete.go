package routes

import (
	"net/http"

	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

// "/edit" route, accessible by only javascript
// updates the document with input form data

func DeleteT(c echo.Context) error {
	id := c.QueryParam("id")

	err := utils.DeleteTransaction(id)
	if err != nil {
		res := utils.CreateResponseMessage(http.StatusBadRequest, false, "Operation failed")
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.CreateResponseMessage(http.StatusOK, true, "Success")
	return c.JSON(http.StatusOK, res)
}
