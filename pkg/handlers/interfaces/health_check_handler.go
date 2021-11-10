package interfaces

import "github.com/labstack/echo/v4"

type HealthCheckHandler interface {
	Live(echo.Context) error
	Ready(echo.Context) error
}
