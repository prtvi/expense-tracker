package main

import (
	"fmt"
	"net/http"

	config "webdev/config"
	model "webdev/model"
	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config.EstablishConnection()

	e := echo.New()
	e.Renderer = model.T
	e.Static("/public", "public")

	e.GET("/", Index)
	e.GET("/add", InputData)

	e.Logger.Fatal(e.Start(":1323"))
}

func Index(c echo.Context) error {
	fmt.Println("hit: GET: /")

	allTransactions := utils.GetTransactions()
	summary := utils.UpdateSummary(allTransactions)

	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"totalExpense":     summary.TotalExpense,
		"totalIncome":      summary.TotalIncome,
		"currentBalance":   summary.CurrentBalance,
		"transactions":     allTransactions,
		"ifNoTransactions": len(allTransactions) == 0,
	})
}

func InputData(c echo.Context) error {
	fmt.Println("hit: GET: /add")

	var message model.ResponseMsg
	transaction, err := utils.InitTransaction(c)
	if err != nil {
		message.StatusCode = http.StatusBadRequest
		message.Success = false

		return c.JSON(http.StatusBadRequest, message)
	}

	err = utils.InsertTransaction(transaction)
	if err != nil {
		message.StatusCode = http.StatusBadRequest
		message.Success = false

		return c.JSON(http.StatusBadRequest, message)
	}

	message.StatusCode = http.StatusOK
	message.Success = true

	return c.JSON(http.StatusOK, message)
}
