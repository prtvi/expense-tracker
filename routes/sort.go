package routes

import (
	"net/http"

	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

// "/sort" route, accessible by only javascript

func Sort(c echo.Context) error {
	res := utils.CreateResponseMessage(http.StatusOK, true, "Success")
	return c.JSON(http.StatusOK, res)
}
