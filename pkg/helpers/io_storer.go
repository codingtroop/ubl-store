package helpers

import (
	"context"
	"os"

	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

type ioStorer struct {
	folder string
}

func NewIOStorer(f string) interfaces.Storer {
	return &ioStorer{folder: f}
}
func (h *ioStorer) Read(c context.Context, uuid string) ([]byte, error) {
	return os.ReadFile(h.folder + "/" + uuid)
}

func (h *ioStorer) Write(c context.Context, uuid string, data []byte) error {
	if _, err := os.Stat(h.folder); os.IsNotExist(err) {
		os.MkdirAll(h.folder, os.ModePerm)
	}
	return os.WriteFile(h.folder+"/"+uuid, data, 0644)
}
