package main

import (
	"net/http"

	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/hello/:name", func(c echo.Context) error {
		fmt.Println(c.FormValue("age"))
		return c.JSON(http.StatusOK, map[string]string{
			"name": c.Param("name"),
			"age":  c.QueryParam("age"),
		})
	})
	e.GET("/world", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>go的服务器</h1>")
	})
	e.Start(":1323")
}
