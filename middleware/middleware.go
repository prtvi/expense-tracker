package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		dt := time.Now().String()[0:19]
		fmt.Printf("%s: %s on %s\n\n", c.Request().Method, c.Request().URL, dt)
		return next(c)
	}
}
