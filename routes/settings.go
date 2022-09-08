package routes

import (
	utils "github.com/prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// "/settings" route to accept user settings

func Settings(c echo.Context) error {
	settings := utils.ParseSettingsUrlParams(c)
	res := utils.GetResponseMessage(false)

	// init/update user settings
	err := utils.InitAndUpdateSettings(settings)
	if err != nil {
		return c.JSON(res.StatusCode, res)
	}

	// init user budget
	err = utils.InitAndSetBudget(float32(settings.CurrentMonthBudget))
	if err != nil {
		return c.JSON(res.StatusCode, res)
	}

	// init user summary
	err = utils.InitSummary()
	if err != nil {
		return c.JSON(res.StatusCode, res)
	}

	res = utils.GetResponseMessage(true)
	return c.JSON(res.StatusCode, res)
}
