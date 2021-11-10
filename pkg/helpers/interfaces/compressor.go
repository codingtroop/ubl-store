package interfaces

import "context"

type Compressor interface {
	Compress(c context.Context, fileName string, data []byte) ([]byte, error)
	Decompress(c context.Context, data []byte) ([]byte, error)
}
