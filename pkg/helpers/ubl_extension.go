package helpers

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"

	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
	"github.com/google/uuid"
)

type ublExtension struct {
}

func NewUblExtension() interfaces.UblExtension {
	return &ublExtension{}
}

func (u *ublExtension) Hash(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func (u *ublExtension) ParseUbl(data []byte) (string, string, *map[string]string, error) {
	bs := string(data)

	su := "<cbc:UUID>"
	eu := "</cbc:UUID>"

	sui := strings.Index(bs, su)
	eui := strings.Index(bs, eu)

	id := bs[sui+len(su) : eui]

	st := "<cac:AdditionalDocumentReference>"
	et := "</cac:AdditionalDocumentReference>"
	bo := "</cbc:EmbeddedDocumentBinaryObject>"

	sti := strings.Index(bs, st)

	if sti == -1 {
		return bs, id, nil, nil
	}

	eti := strings.LastIndex(bs, et)

	t := bs[sti : eti+len(et)]

	ta := strings.Split(t, st)

	tam := map[string]string{}

	for _, v := range ta {

		if v == "" {
			continue
		}

		tuuid := uuid.New().String()

		r := regexp.MustCompile(`EmbeddedDocumentBinaryObject\b.*\w*>\b`)

		f := r.FindString(v)

		v := v[strings.Index(v, f)+len(f) : strings.Index(v, bo)]

		bs = strings.ReplaceAll(bs, v, tuuid)

		tam[tuuid] = v
	}

	return bs, id, &tam, nil
}
