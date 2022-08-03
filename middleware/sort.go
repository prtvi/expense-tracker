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
		now := time.Now()
		currentYear, currentMonth, currentDay := now.Date()
		currentLocation := now.Location()

		// endDate initially initialized with hour, minute and second as 0s
		var viewEndDate time.Time = time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)

		// start from date (past)
		var viewStartDate time.Time

		// if view is empty, then set it to config.ViewAll
		if view == "" {
			view = config.ViewAll
		}

		switch view {
		// last 7 days
		case config.ViewLast7Days:
			viewStartDate = viewEndDate.AddDate(0, 0, -7+1)

		// last 30 days
		case config.ViewLast30Days:
			viewStartDate = viewEndDate.AddDate(0, 0, -30+1)

		// this month
		case config.ViewThisMonth:
			firstOfMonth, lastOfMonth := utils.FirstAndLastDayOfMonth(currentYear, int(currentMonth), currentLocation)

			viewStartDate = firstOfMonth
			viewEndDate = lastOfMonth

		// last month
		case config.ViewLastMonth:
			firstOfMonth, lastOfMonth := utils.FirstAndLastDayOfMonth(currentYear, int(currentMonth-1), currentLocation)

			viewStartDate = firstOfMonth
			viewEndDate = lastOfMonth

		// custom
		case config.ViewCustom:
			dateStart, dateEnd := c.QueryParam(config.CustomDateStartID), c.QueryParam(config.CustomDateEndID)

			viewStartDate = utils.DateStringToDateObj(dateStart, false)
			viewEndDate = utils.DateStringToDateObj(dateEnd, false)

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

		// adding a day to the endDate and subtracting a nanosecond to make the endDate time to 23:59:59.9999 rather than 00:00:00
		// for sorting dates "lesser than equal to" endDate

		if view != config.ViewAll {
			viewEndDate = utils.LastSecondOfTheDay(viewEndDate)
		}

		c.Set(config.View, view)
		c.Set(config.ViewStartDate, viewStartDate)
		c.Set(config.ViewEndDate, viewEndDate)

		c.Set(config.Sort, sort)

		return next(c)
	}
}
