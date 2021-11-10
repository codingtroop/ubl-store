package implementations

import (
	"net/http"

	"github.com/codingtroop/ubl-store/pkg/handlers/interfaces"
	"github.com/labstack/echo/v4"
)

type healthCheckHandler struct {
}

func NewHealthCheckHandler() interfaces.HealthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) Live(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func (h *healthCheckHandler) Ready(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
