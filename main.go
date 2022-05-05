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
	e.POST("/add", InputData)

	e.Logger.Fatal(e.Start(":1323"))
}

func Index(c echo.Context) error {
	fmt.Println("hit: GET: /")

	allTransactions := utils.GetTransactions()
	currentStatus := utils.UpdateCurrentStatus(allTransactions)

	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"totalExpense":   currentStatus.TotalExpense,
		"totalIncome":    currentStatus.TotalIncome,
		"currentBalance": currentStatus.CurrentBalance,
		"transactions":   allTransactions,
	})
}

func InputData(c echo.Context) error {
	fmt.Println("hit: POST: /add")

	transaction := utils.InitTransaction(c)
	utils.InsertTransaction(transaction)
	allTransactions := utils.GetTransactions()
	currentStatus := utils.UpdateCurrentStatus(allTransactions)

	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"totalExpense":   currentStatus.TotalExpense,
		"totalIncome":    currentStatus.TotalIncome,
		"currentBalance": currentStatus.CurrentBalance,
		"transactions":   allTransactions,
	})
}
