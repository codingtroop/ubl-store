package main

import (
	"net/http"

	_ "github.com/codingtroop/ubl-store/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hellos, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
