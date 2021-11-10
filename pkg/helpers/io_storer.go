package helpers

import (
	"context"
	"os"

	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

type ioStorer struct {
}

var root = "var/ubl"

func NewIOHelper() interfaces.Storer {
	return &ioStorer{}
}
func (h *ioStorer) Read(c context.Context, uuid string) ([]byte, error) {
	return os.ReadFile(root + "/" + uuid)
}

func (h *ioStorer) Write(c context.Context, uuid string, data []byte) error {
	return os.WriteFile(root+"/"+uuid, data, 0644)
}
