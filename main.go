package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleIndex(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "Hello World, "+name)
}

func main() {
	e := echo.New()

	// ===== before refactor
	// e.GET("/", func(c echo.Context) error {
	//     return c.String(http.StatusOK, "Hello, World!")
	// })

	// ===== after refactor
	e.GET("/", HandleIndex)
	e.POST("/", HandleIndex)

	e.Logger.Fatal(e.Start(":1323"))
}
