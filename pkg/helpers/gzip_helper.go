package helpers

import (
	"bytes"
	"compress/gzip"
	"context"
	"io/ioutil"

	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

type gzipHelper struct {
}

func NewGZipHelper() interfaces.ZipHelper {
	return &gzipHelper{}
}

func (h *gzipHelper) Zip(c context.Context, fileName string, data []byte) ([]byte, error) {
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

func (h *gzipHelper) Unzip(c context.Context, data []byte) (string, []byte, error) {
	zr, err := gzip.NewReader(bytes.NewReader(data))

	if err != nil {
		return "", nil, err
	}

	d, err := ioutil.ReadAll(zr)

	if err != nil {
		return "", nil, err
	}

	if err := zr.Close(); err != nil {
		return "", nil, err
	}

	return zr.Name, d, nil

}
