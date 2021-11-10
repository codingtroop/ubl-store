package interfaces

import "context"

type FileHelper interface {
	Read(c context.Context, filePath string) ([]byte, error)
	Write(c context.Context, filePath string, data []byte) error
}
