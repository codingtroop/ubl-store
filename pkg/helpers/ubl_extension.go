package helpers

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

type ublExtension struct {
}

func NewUblExtension() interfaces.UblExtension {
	return &ublExtension{}
}

func (u *ublExtension) Hash(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func (u *ublExtension) GetCertificate(data []byte) (string, error) {

	doc, err := xmlquery.Parse(bytes.NewReader(data))

	if err != nil {
		return "", err
	}
	certNode := xmlquery.FindOne(doc, "//*[local-name()='X509Data']")

	if certNode == nil {
		return "", nil
	}

	return certNode.OutputXML(true), nil
}

func (u *ublExtension) GetAdditionalInfo(data []byte) (*[]string, error) {

	doc, err := xmlquery.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var prefix string

	if uuidNode := xmlquery.FindOne(doc, "//*[local-name()='UUID']"); uuidNode != nil {
		prefix = uuidNode.Prefix
	}

	bo := "</" + prefix + ":EmbeddedDocumentBinaryObject>"

	attNodes := xmlquery.Find(doc, "//*[local-name()='AdditionalDocumentReference']")

	certNode := xmlquery.FindOne(doc, "//*[local-name()='X509Data']")

	if len(attNodes) == 0 && certNode == nil {
		return nil, nil
	}

	tam := []string{}

	if certNode != nil {
		tam = append(tam, certNode.OutputXML(false))
	}

	for _, n := range attNodes {

		v := n.OutputXML(true)

		r := regexp.MustCompile(`EmbeddedDocumentBinaryObject\b.*\w*>\b`)

		f := r.FindString(v)

		v = v[strings.Index(v, f)+len(f) : strings.Index(v, bo)]

		tam = append(tam, v)
	}

	return &tam, nil
}

func (u *ublExtension) Parse(data []byte) (string, string, *map[string]string, error) {
	bs := string(data)

	doc, err := xmlquery.Parse(bytes.NewReader(data))
	if err != nil {
		return "", "", nil, err
	}

	var id string
	var prefix string

	if uuidNode := xmlquery.FindOne(doc, "//*[local-name()='UUID']"); uuidNode != nil {
		id = uuidNode.InnerText()
		prefix = uuidNode.Prefix
	}

	bo := "</" + prefix + ":EmbeddedDocumentBinaryObject>"

	attNodes := xmlquery.Find(doc, "//*[local-name()='AdditionalDocumentReference']")

	certNode := xmlquery.FindOne(doc, "//*[local-name()='X509Data']")

	if len(attNodes) == 0 && certNode == nil {
		return bs, id, nil, nil
	}

	tam := map[string]string{}

	if certNode != nil {
		xds := fmt.Sprintf("<%s:X509Data>", certNode.Prefix)
		xde := fmt.Sprintf("</%s:X509Data>", certNode.Prefix)

		xsi := strings.Index(bs, xds)
		xei := strings.Index(bs, xde)

		cert := bs[xsi+len(xds) : xei]
		certHash := u.Hash(cert)
		bs = strings.ReplaceAll(bs, cert, certHash)

		tam[certHash] = cert
	}

	for _, n := range attNodes {

		v := n.OutputXML(true)

		r := regexp.MustCompile(`EmbeddedDocumentBinaryObject\b.*\w*>\b`)

		f := r.FindString(v)

		v = v[strings.Index(v, f)+len(f) : strings.Index(v, bo)]

		h := u.Hash(v)

		bs = strings.ReplaceAll(bs, v, h)

		tam[h] = v
	}

	return bs, id, &tam, nil
}
