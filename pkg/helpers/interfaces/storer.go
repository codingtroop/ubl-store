package interfaces

import "context"

type Storer interface {
	Exists(context.Context, string) (bool, error)
	Read(c context.Context, uuid string) ([]byte, error)
	Write(c context.Context, uuid string, data []byte) error
}
