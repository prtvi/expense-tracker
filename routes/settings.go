package routes

import (
	"net/http"
	"prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/settings" route to accept user settings

func Settings(c echo.Context) error {
	settings := utils.ParseSettingsUrlParams(c)

	// init/update user settings
	err := utils.InitAndUpdateSettings(settings)
	if err != nil {
		res := utils.CreateResponseMessage(http.StatusBadRequest, false, "Operation failed")
		return c.JSON(http.StatusBadRequest, res)
	}

	// init user budget
	err = utils.InitAndSetBudget(float32(settings.CurrentMonthBudget))
	if err != nil {
		res := utils.CreateResponseMessage(http.StatusBadRequest, false, "Operation failed")
		return c.JSON(http.StatusBadRequest, res)
	}

	// init user summary
	err = utils.InitSummary()
	if err != nil {
		res := utils.CreateResponseMessage(http.StatusBadRequest, false, "Operation failed")
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.CreateResponseMessage(http.StatusOK, true, "Success")
	return c.JSON(http.StatusOK, res)
}
