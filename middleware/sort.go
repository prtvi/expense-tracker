package middleware

import (
	"time"

	config "webdev/config"
	"webdev/utils"

	"github.com/labstack/echo/v4"
)

// will set the range of dates to sort from
func Sort(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sortParam := c.QueryParam(config.SortInputID)

		// end date (today date)
		var sortEndDate time.Time = time.Now()
		// start from date (past)
		var sortStartDate time.Time

		switch sortParam {
		// 7
		case config.SortLast7DaysValue:
			sortStartDate = sortEndDate.AddDate(0, 0, -7)

		// 30
		case config.SortLast30DaysValue:
			sortStartDate = sortEndDate.AddDate(0, 0, -30)

		// this month
		case config.SortThisMonthValue:
			now := time.Now()
			currentYear, currentMonth, _ := now.Date()
			currentLocation := now.Location()

			firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
			lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

			sortStartDate = firstOfMonth
			sortEndDate = lastOfMonth

		// last month
		case config.SortLastMonthValue:
			now := time.Now()
			currentYear, currentMonth, _ := now.Date()
			currentLocation := now.Location()

			firstOfMonth := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, currentLocation)
			lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

			sortStartDate = firstOfMonth
			sortEndDate = lastOfMonth

		// custom
		case config.SortCustomValue:
			dateStart, dateEnd := c.QueryParam(config.CustomDateStartID), c.QueryParam(config.CustomDateEndID)

			sortStartDate = utils.DateStringToDateObj(dateStart)
			sortEndDate = utils.DateStringToDateObj(dateEnd)

		// all
		case config.SortAllValue:
			fallthrough

		default:
			sortStartDate, sortEndDate = utils.GetNewestAndOldestTDates()
		}

		c.Set(config.SortEndDate, sortEndDate)
		c.Set(config.SortStartDate, sortStartDate)

		return next(c)
	}
}
