package routes

import (
	"net/http"
	"prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/settings" route to accept user settings

func Settings(c echo.Context) error {
	settings := utils.ParseSettingsUrlParams(c)

	success := utils.UpdateSettings(settings)
	if !success {
		res := utils.CreateResponseMessage(http.StatusBadRequest, false, "Operation failed")
		return c.JSON(http.StatusBadRequest, res)
	}

	utils.SetBudget(float32(settings.CurrentMonthBudget))

	res := utils.CreateResponseMessage(http.StatusOK, true, "Success")
	return c.JSON(http.StatusOK, res)
}
