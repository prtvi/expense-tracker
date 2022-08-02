package main

import (
	config "prtvi/expense-tracker/config"
	middleware "prtvi/expense-tracker/middleware"
	model "prtvi/expense-tracker/model"
	routes "prtvi/expense-tracker/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.EstablishConnection()

	e := echo.New()
	e.Renderer = model.T
	e.Static("/public", "public")

	e.GET("/", routes.Home, middleware.Logger, middleware.Sort)
	e.GET("/get", routes.ReturnT, middleware.Logger)
	e.GET("/add", routes.AddT, middleware.Logger)
	e.GET("/edit", routes.EditT, middleware.Logger)
	e.GET("/del", routes.DeleteT, middleware.Logger)

	e.Logger.Fatal(e.Start(":1323"))
}
