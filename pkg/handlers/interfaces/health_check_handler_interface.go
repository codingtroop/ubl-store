package interfaces

import "github.com/labstack/echo/v4"

type HealthCheckHandlerInterface interface {
	Live(echo.Context) error
	Ready(echo.Context) error
}
