package routes

import (
	"net/http"

	config "github.com/prtvi/expense-tracker/config"
	utils "github.com/prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

func Summary(c echo.Context) error {
	start, end := c.QueryParam("start"), c.QueryParam("end")

	startTime := utils.DateStringToDateObj(start)
	endTime := utils.DateStringToDateObj(end)
	endTime = utils.GoToLastSecondOfTheDay(endTime)

	summary := utils.GetSummary(startTime, endTime, config.SortAscID)

	return c.JSON(http.StatusOK, summary)
}
