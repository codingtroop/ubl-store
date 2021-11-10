package helpers

import (
	"context"
	"os"

	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

type ioHelper struct {
}

func NewIOHelper() interfaces.FileHelper {
	return &ioHelper{}
}

func (h *ioHelper) Read(c context.Context, filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func (h *ioHelper) Write(c context.Context, filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}
