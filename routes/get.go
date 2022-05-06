package routes

import (
	"fmt"
	"net/http"

	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

// "/get" route, accessible by only javascript
// returns a transaction by doc id

func ReturnT(c echo.Context) error {
	fmt.Println("hit: GET: /get")

	id := c.QueryParam("id")

	transaction, err := utils.GetDocumentById(id)
	if err != nil {
		message := utils.CreateResponseMessage(http.StatusBadRequest, false)
		return c.JSON(http.StatusBadRequest, message)
	}

	return c.JSON(http.StatusOK, transaction)
}
