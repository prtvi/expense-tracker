package middleware

import (
	"time"

	config "webdev/config"
	utils "webdev/utils"

	"github.com/labstack/echo/v4"
)

// will set the range of dates to sort from
func Sort(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		filterBy := c.QueryParam(config.FilterBy)
		sortBy := c.QueryParam(config.SortBy)

		// only for not "custom" options
		// end date (today date)
		var sortEndDate time.Time = time.Now()
		// start from date (past)
		var sortStartDate time.Time

		switch filterBy {
		// 7
		case config.FilterLast7Days:
			sortStartDate = sortEndDate.AddDate(0, 0, -7)

		// 30
		case config.FilterLast30Days:
			sortStartDate = sortEndDate.AddDate(0, 0, -30)

		// this month
		case config.FilterThisMonth:
			now := time.Now()
			currentYear, currentMonth, _ := now.Date()
			currentLocation := now.Location()

			firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
			lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

			sortStartDate = firstOfMonth
			sortEndDate = lastOfMonth

		// last month
		case config.FilterLastMonth:
			now := time.Now()
			currentYear, currentMonth, _ := now.Date()
			currentLocation := now.Location()

			firstOfMonth := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, currentLocation)
			lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

			sortStartDate = firstOfMonth
			sortEndDate = lastOfMonth

		// custom
		case config.FilterCustom:
			dateStart, dateEnd := c.QueryParam(config.CustomDateStartID), c.QueryParam(config.CustomDateEndID)

			sortStartDate = utils.DateStringToDateObj(dateStart)
			sortEndDate = utils.DateStringToDateObj(dateEnd)

		// all
		case config.FilterAll:
			fallthrough

		default:
			var err error
			sortStartDate, sortEndDate, err = utils.GetNewestAndOldestTDates()
			if err != nil {
				break
			}
		}

		// if filterBy is empty, then set it to config.FilterAll
		if filterBy == "" {
			filterBy = config.FilterAll
		}

		c.Set(config.FilterBy, filterBy)
		c.Set(config.SortBy, sortBy)
		c.Set(config.SortEndDate, sortEndDate)
		c.Set(config.SortStartDate, sortStartDate)

		return next(c)
	}
}
