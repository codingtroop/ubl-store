package handlers

import (
	"net/http"

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

// Get godoc
// @Summary Get ubl
// @Tags Ubl
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200
// @Router /api/v1/ubl/{id} [get]
func (h *ublStoreHandler) Get(c echo.Context) error {

	model := models.GetModel{}

	if err := c.Bind(model); err != nil {
		return err
	}

	context := c.Request().Context()

	ubl, err := h.ublRepo.Get(context, model.ID)

	if err != nil {
		return err
	}

	if ubl == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, ubl)
}

// Post godoc
// @Summary Post ubl
// @Tags Ubl
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/ubl [post]
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

// Delete godoc
// @Summary Delete ubl
// @Tags Ubl
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200
// @Router /api/v1/ubl/{id} [delete]
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
