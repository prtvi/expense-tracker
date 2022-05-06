package main

import (
	config "webdev/config"
	model "webdev/model"
	routes "webdev/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.EstablishConnection()

	e := echo.New()
	e.Renderer = model.T
	e.Static("/public", "public")

	e.GET("/", routes.Home)
	e.GET("/get", routes.ReturnT)
	e.GET("/add", routes.AddT)
	e.GET("/edit", routes.EditT)

	e.Logger.Fatal(e.Start(":1324"))
}
