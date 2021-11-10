package helpers

import (
	"bytes"
	"compress/gzip"
	"context"
	"io/ioutil"

	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

type gzipper struct {
}

func NewGZip() interfaces.Compressor {
	return &gzipper{}
}

func (h *gzipper) Compress(c context.Context, fileName string, data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	zw.Name = fileName

	if _, err := zw.Write(data); err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (h *gzipper) Decompress(c context.Context, data []byte) ([]byte, error) {
	zr, err := gzip.NewReader(bytes.NewReader(data))

	if err != nil {
		return nil, err
	}

	d, err := ioutil.ReadAll(zr)

	if err != nil {
		return nil, err
	}

	if err := zr.Close(); err != nil {
		return nil, err
	}

	return d, nil

}
