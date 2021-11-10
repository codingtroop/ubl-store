package helpers

import (
	"archive/zip"
	"bytes"
	"context"
	"io/ioutil"

	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

type zipHelper struct {
}

func NewZipHelper() interfaces.ZipHelper {
	return &zipHelper{}
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func (h *zipHelper) Zip(c context.Context, data map[string][]byte) ([]byte, error) {
	buf := new(bytes.Buffer)

	zw := zip.NewWriter(buf)

	for f, d := range data {
		zf, err := zw.Create(f)
		if err != nil {
			return nil, err
		}
		_, err = zf.Write([]byte(d))
		if err != nil {
			return nil, err
		}
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (h *zipHelper) Unzip(c context.Context, data []byte) (map[string][]byte, error) {
	files := map[string][]byte{}
	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))

	if err != nil {
		return nil, err
	}

	for _, z := range zr.File {
		d, err := readZipFile(z)
		if err != nil {
			return nil, err
		}

		files[z.Name] = d
	}
	return files, nil
}
