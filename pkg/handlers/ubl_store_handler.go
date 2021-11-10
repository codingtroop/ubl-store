package handlers

import (
	"net/http"
	"strings"

	"github.com/codingtroop/ubl-store/pkg/entities"
	handler "github.com/codingtroop/ubl-store/pkg/handlers/interfaces"
	"github.com/codingtroop/ubl-store/pkg/models"
	repo "github.com/codingtroop/ubl-store/pkg/repositories/interfaces"
	"github.com/labstack/echo/v4"
)

type ublStoreHandler struct {
	ublRepo        repo.UblRepository
	attachmentRepo repo.AttachmentRepository
}

func NewUblStoreHandler(ur repo.UblRepository,
	ar repo.AttachmentRepository) handler.UblStoreHandler {
	return &ublStoreHandler{ublRepo: ur, attachmentRepo: ar}
}

func (h *ublStoreHandler) Get(c echo.Context) error {

	model := models.GetModel{}

	if err := c.Bind(model); err != nil {
		return err
	}

	context := c.Request().Context()

	ubl, err := h.ublRepo.Get(context, model.ID)

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return c.NoContent(http.StatusNotFound)
		}

		return err
	}

	return c.JSON(http.StatusOK, ubl)
}

func (h *ublStoreHandler) Post(c echo.Context) error {

	model := models.PostModel{}

	if err := c.Bind(model); err != nil {
		return err
	}

	ubl := entities.UblEntity{}

	context := c.Request().Context()

	if err := h.ublRepo.Insert(context, ubl); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *ublStoreHandler) Delete(c echo.Context) error {

	model := models.DeleteModel{}

	if err := c.Bind(model); err != nil {
		return err
	}

	context := c.Request().Context()

	if err := h.ublRepo.Delete(context, model.ID); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
