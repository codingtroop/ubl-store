package handlers

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/codingtroop/ubl-store/pkg/entities"
	handler "github.com/codingtroop/ubl-store/pkg/handlers/interfaces"
	helpers "github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
	"github.com/codingtroop/ubl-store/pkg/models"
	repo "github.com/codingtroop/ubl-store/pkg/repositories/interfaces"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ublStoreHandler struct {
	ublRepo         repo.UblRepository
	attachmentRepo  repo.AttachmentRepository
	ublStore        helpers.Storer
	attachmentStore helpers.Storer
	compressor      helpers.Compressor
	ubl             helpers.UblExtension
}

func NewUblStoreHandler(ur repo.UblRepository,
	ar repo.AttachmentRepository, us helpers.Storer, as helpers.Storer, c helpers.Compressor, u helpers.UblExtension) handler.UblStoreHandler {
	return &ublStoreHandler{ublRepo: ur, attachmentRepo: ar, ublStore: us, attachmentStore: as, compressor: c, ubl: u}
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

	id := c.Param("id")

	guid, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	context := c.Request().Context()

	ubl, err := h.ublRepo.Get(context, guid)

	if err != nil {
		return err
	}

	if ubl == nil {
		return c.NoContent(http.StatusNotFound)
	}

	zipData, err := h.ublStore.Read(context, id)

	if err != nil {
		return err
	}

	if zipData == nil {
		return c.NoContent(http.StatusNotFound)
	}

	data, err := h.compressor.Decompress(context, zipData)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &models.UblModel{Data: data})
}

// Post godoc
// @Summary Post ubl
// @Tags Ubl
// @Accept  json
// @Produce  json
// @Success 200
// @Param data body models.PostModel true "data"
// @Router /api/v1/ubl [post]
func (h *ublStoreHandler) Post(c echo.Context) error {

	model := models.PostModel{}

	if err := c.Bind(&model); err != nil {
		return err
	}

	b, err := base64.StdEncoding.DecodeString(model.Data)

	if err != nil {
		return err
	}

	ublText, uuidText, attach, err := h.ubl.ParseUbl(b)

	if err != nil {
		return err
	}

	ubl := entities.UblEntity{Created: time.Now()}

	if id, err := uuid.Parse(uuidText); err != nil {
		return err
	} else {
		ubl.ID = id
	}

	context := c.Request().Context()

	cBytes, err := h.compressor.Compress(context, ublText, b)

	if err != nil {
		return err
	}

	if err := h.ublStore.Write(context, uuidText, cBytes); err != nil {
		return err
	}

	if err := h.ublRepo.Insert(context, ubl); err != nil {
		return err
	}

	if attach != nil {
		for aid, v := range *attach {

			att := entities.AttachmentEntity{UblID: ubl.ID, Created: time.Now(), Hash: h.ubl.Hash(v)}

			if id, err := uuid.Parse(aid); err != nil {
				return err
			} else {
				att.ID = id
			}

			cBytes, err := h.compressor.Compress(context, v, b)

			if err != nil {
				return err
			}

			if err := h.ublStore.Write(context, aid, cBytes); err != nil {
				return err
			}

			if err := h.attachmentRepo.Insert(context, att); err != nil {
				return err
			}
		}
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
