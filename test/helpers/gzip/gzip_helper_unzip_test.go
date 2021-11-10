package test

import (
	"context"
	"testing"

	"github.com/codingtroop/ubl-store/pkg/helpers"
)

func Test_GZip_UnZip_Ok(t *testing.T) {
	zh := helpers.NewGZip()

	d, err := zh.Compress(context.TODO(), "test", []byte("test"))

	if err != nil {
		t.Fatal(err)
	}

	if _, err := zh.Decompress(context.TODO(), d); err != nil {
		t.Fatal(err)
	}
}
