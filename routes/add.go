package routes

import (
	"fmt"
	"net/http"

	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

// "/add" route, accessible by only javascript
// adds a new transaction to the db

func AddT(c echo.Context) error {
	fmt.Println("hit: GET: /add")

	transaction, err := utils.InitTransaction(c)
	if err != nil {
		message := utils.CreateResponseMessage(http.StatusBadRequest, false)
		return c.JSON(http.StatusBadRequest, message)
	}

	err = utils.InsertTransaction(transaction)
	if err != nil {
		message := utils.CreateResponseMessage(http.StatusBadRequest, false)
		return c.JSON(http.StatusBadRequest, message)
	}

	message := utils.CreateResponseMessage(http.StatusOK, true)
	return c.JSON(http.StatusOK, message)
}
