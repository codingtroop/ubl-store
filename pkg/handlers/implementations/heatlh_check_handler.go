package implementations

import (
	"net/http"

	"github.com/labstack/echo"
)

type healthCheckHandler struct {
}

func NewHealthCheckHandler() HealthCheckHandlerInterface {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) Live(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func (h *healthCheckHandler) Ready(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
