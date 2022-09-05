package routes

import (
	"net/http"
	"prtvi/expense-tracker/config"
	utils "prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

func Summary(c echo.Context) error {
	start, end := c.QueryParam("start"), c.QueryParam("end")

	startTime := utils.DateStringToDateObj(start, true)
	endTime := utils.DateStringToDateObj(end, true)

	summary := utils.GetSummary(startTime, endTime, config.SortAscID)

	return c.JSON(http.StatusOK, summary)
}
