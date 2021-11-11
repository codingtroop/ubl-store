package helpers

import (
	"bytes"

	"github.com/antchfx/xmlquery"
	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

type ublExtension struct {
}

func NewUblExtension() interfaces.UblExtension {
	return &ublExtension{}
}

func (u *ublExtension) GetUUID(data []byte) (string, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(data))

	if err != nil {
		return "", err
	}

	uuidNode := xmlquery.FindOne(doc, "//cbc:UUID")

	return uuidNode.InnerText(), nil
}
