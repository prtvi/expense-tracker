package main

import (
	cfg "github.com/prtvi/expense-tracker/config"
	mw "github.com/prtvi/expense-tracker/middleware"
	uModel "github.com/prtvi/expense-tracker/model/utils"
	r "github.com/prtvi/expense-tracker/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg.EstablishConnection()

	e := echo.New()
	e.Renderer = uModel.T
	e.Static("/public", "public")

	e.GET("/", r.Home, mw.Logger, mw.Sort)

	// routes for javascript to make requests
	e.GET("/get", r.ReturnT, mw.Logger)
	e.GET("/add", r.AddT, mw.Logger)
	e.GET("/edit", r.EditT, mw.Logger)
	e.GET("/del", r.DeleteT, mw.Logger)
	e.GET("/settings", r.Settings, mw.Logger)
	e.GET("/summary", r.Summary, mw.Logger)

	e.Logger.Fatal(e.Start(":1323"))
}
