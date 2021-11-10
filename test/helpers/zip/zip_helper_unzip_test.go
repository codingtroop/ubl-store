package test

import (
	"context"
	"testing"

	"github.com/codingtroop/ubl-store/pkg/helpers"
)

func Test_Zip_UnZip_Ok(t *testing.T) {
	zh := helpers.NewZipper()

	d, err := zh.Compress(context.TODO(), "test", []byte("test"))

	if err != nil {
		t.Fatal(err)
	}

	if _, _, err := zh.Decompress(context.TODO(), d); err != nil {
		t.Fatal(err)
	}
}
