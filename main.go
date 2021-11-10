package main

import (
	"net/http"

	_ "github.com/codingtroop/ubl-store/docs"
	api "github.com/codingtroop/ubl-store/pkg/handlers/implementations"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	hc := api.NewHealthCheckHandler()

	e.GET("/health", hc.Live)
	e.GET("/health/live", hc.Live)
	e.GET("/health/ready", hc.Ready)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World haha oldu!")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
