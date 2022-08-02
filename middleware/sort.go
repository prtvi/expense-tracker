package middleware

import (
	"time"

	config "prtvi/expense-tracker/config"
	utils "prtvi/expense-tracker/utils"

	"github.com/labstack/echo/v4"
)

// will set the range of dates to sort from
func Sort(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		view := c.QueryParam(config.View)
		sort := c.QueryParam(config.Sort)

		// only for not "custom" options
		// end date (today date)
		var viewEndDate time.Time = time.Now()

		// start from date (past)
		var viewStartDate time.Time

		switch view {
		// 7
		case config.ViewLast7Days:
			viewStartDate = viewEndDate.AddDate(0, 0, -7)

		// 30
		case config.ViewLast30Days:
			viewStartDate = viewEndDate.AddDate(0, 0, -30)

		// this month
		case config.ViewThisMonth:
			now := time.Now()
			currentYear, currentMonth, _ := now.Date()
			currentLocation := now.Location()

			firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
			lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

			viewStartDate = firstOfMonth
			viewEndDate = lastOfMonth

		// last month
		case config.ViewLastMonth:
			now := time.Now()
			currentYear, currentMonth, _ := now.Date()
			currentLocation := now.Location()

			firstOfMonth := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, currentLocation)
			lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

			viewStartDate = firstOfMonth
			viewEndDate = lastOfMonth

		// custom
		case config.ViewCustom:
			dateStart, dateEnd := c.QueryParam(config.CustomDateStartID), c.QueryParam(config.CustomDateEndID)

			viewStartDate = utils.DateStringToDateObj(dateStart)
			viewEndDate = utils.DateStringToDateObj(dateEnd)

		// all
		case config.ViewAll:
			fallthrough

		default:
			var err error
			viewStartDate, viewEndDate, err = utils.GetNewestAndOldestTDates()
			if err != nil {
				break
			}
		}

		// if view is empty, then set it to config.ViewAll
		if view == "" {
			view = config.ViewAll
		}

		c.Set(config.View, view)
		c.Set(config.ViewStartDate, viewStartDate)
		c.Set(config.ViewEndDate, viewEndDate)

		c.Set(config.Sort, sort)

		return next(c)
	}
}
