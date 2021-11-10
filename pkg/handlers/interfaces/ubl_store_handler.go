package interfaces

import (
	"github.com/labstack/echo/v4"
)

type UblStoreHandler interface {
	Get(echo.Context) error
	Post(echo.Context) error
	Delete(echo.Context) error
}
