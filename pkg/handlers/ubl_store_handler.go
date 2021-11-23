package handlers

import (
	"encoding/base64"
	"net/http"
	"strings"

	handler "github.com/codingtroop/ubl-store/pkg/handlers/interfaces"
	helpers "github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
	"github.com/codingtroop/ubl-store/pkg/models"
	"github.com/labstack/echo/v4"
)

type ublStoreHandler struct {
	ublStore        helpers.Storer
	attachmentStore helpers.Storer
	compressor      helpers.Compressor
	ubl             helpers.UblExtension
}

func NewUblStoreHandler(us helpers.Storer, as helpers.Storer, c helpers.Compressor, u helpers.UblExtension) handler.UblStoreHandler {
	return &ublStoreHandler{ublStore: us, attachmentStore: as, compressor: c, ubl: u}
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

	context := c.Request().Context()

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

	atts, err := h.ubl.GetAdditionalDocumentReferances(data)

	if err != nil {
		return err
	}

	sdata := string(data)

	for _, hash := range *atts {
		zipData, err := h.attachmentStore.Read(context, hash)

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

		sdata = strings.ReplaceAll(sdata, hash, string(data))
	}

	return c.JSON(http.StatusOK, &models.UblModel{Data: []byte(sdata)})
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

	ublText, uuidText, attach, err := h.ubl.Parse(b)

	if err != nil {
		return err
	}

	context := c.Request().Context()

	cBytes, err := h.compressor.Compress(context, uuidText, []byte(ublText))

	if err != nil {
		return err
	}

	if err := h.ublStore.Write(context, uuidText, cBytes); err != nil {
		return err
	}

	for hash, v := range *attach {

		//check hash exist on store
		exist, err := h.attachmentStore.Exists(context, hash)

		if err != nil {
			return err
		}

		if !exist {
			cBytes, err := h.compressor.Compress(context, hash, []byte(v))

			if err != nil {
				return err
			}

			if err := h.attachmentStore.Write(context, hash, cBytes); err != nil {
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

	return c.NoContent(http.StatusOK)
}
