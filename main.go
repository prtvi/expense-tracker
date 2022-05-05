// gin --all -i run main.go
package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var t = &Template{
	templates: template.Must(template.ParseGlob("views/*.html")),
}

func main() {
	// config.EstablishConnection()

	e := echo.New()
	e.Renderer = t
	e.Static("/public", "public")

	e.GET("/", Index)

	e.Logger.Fatal(e.Start(":1323"))
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "")

	// return c.Render(http.StatusOK, "index", map[string]interface{}{
	// 	"name":      "World",
	// 	"msg":       "Hi there!",
	// 	"condition": true,
	// 	"array":     []int{1, 2, 3, 4, 5},
	// })
}
