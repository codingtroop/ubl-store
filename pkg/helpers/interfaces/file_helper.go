package interfaces

import "context"

type FileHelper interface {
	Read(cntxt context.Context, path string) ([]byte, error)
	Write(cntxt context.Context, path string, data []byte) error
}
