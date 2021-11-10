package interfaces

import "context"

type ZipHelper interface {
	Zip(c context.Context, data map[string][]byte) ([]byte, error)
	Unzip(c context.Context, data []byte) (map[string][]byte, error)
}
