package helpers

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"io/ioutil"

	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

type zipper struct {
}

func NewZipper() interfaces.Compressor {
	return &zipper{}
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func (h *zipper) Compress(c context.Context, fileName string, data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)

	zw := zip.NewWriter(buf)

	zf, err := zw.Create(fileName)
	if err != nil {
		return nil, err
	}
	_, err = zf.Write([]byte(data))
	if err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (h *zipper) Decompress(c context.Context, data []byte) (string, []byte, error) {

	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))

	if err != nil {
		return "", nil, err
	}

	if len(zr.File) == 0 {
		return "", nil, errors.New("no entry")
	}

	z := zr.File[0]

	d, err := readZipFile(z)
	if err != nil {
		return "", nil, err
	}

	return z.Name, d, nil
}
