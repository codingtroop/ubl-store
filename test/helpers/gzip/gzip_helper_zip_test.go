package test

import (
	"context"
	"testing"

	"github.com/codingtroop/ubl-store/pkg/helpers"
)

func Test_GZip_Zip_Ok(t *testing.T) {
	zh := helpers.NewGZipHelper()

	_, err := zh.Zip(context.TODO(), "test", []byte("test"))

	if err != nil {
		t.Fatal(err)
	}
}
