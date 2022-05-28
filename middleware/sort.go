package middleware

import (
	"fmt"

	config "webdev/config"

	"github.com/labstack/echo/v4"
)

// will set the range of dates to sort from
func Sort(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sortParam := c.QueryParam(config.SortInputID)

		c.Set(config.SortUrlKey, sortParam)

		switch sortParam {
		// all
		case config.SortAllValue:
			fmt.Println(config.SortAllValue)

		// 7
		case config.SortLast7DaysValue:
			fmt.Println(config.SortLast7DaysValue)

		// 30
		case config.SortLast30DaysValue:
			fmt.Println(config.SortLast30DaysValue)

		// this month
		case config.SortThisMonthValue:
			fmt.Println(config.SortThisMonthValue)

		// last month
		case config.SortLastMonthValue:
			fmt.Println(config.SortLastMonthValue)

		// custom
		case config.SortCustomValue:
			fmt.Println(config.SortCustomValue)

			dateStart, dateEnd := c.QueryParam(config.CustomDateStartID), c.QueryParam(config.CustomDateEndID)

			fmt.Println(dateStart, dateEnd)

		default:
			fmt.Println(config.SortAllValue)
		}

		return next(c)
	}
}
