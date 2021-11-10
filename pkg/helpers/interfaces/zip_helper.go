package interfaces

import "context"

type ZipHelper interface {
	Zip(c context.Context, fileName string, data []byte) ([]byte, error)
	Unzip(c context.Context, data []byte) (string, []byte, error)
}
