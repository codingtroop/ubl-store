package test

import (
	"context"
	"testing"

	"github.com/codingtroop/ubl-store/pkg/helpers"
)

func Test_GUnZipOk(t *testing.T) {
	zh := helpers.NewGZipHelper()

	d, err := zh.Zip(context.TODO(), "test", []byte("test"))

	if err != nil {
		t.Fatal(err)
	}

	if _, _, err := zh.Unzip(context.TODO(), d); err != nil {
		t.Fatal(err)
	}
}
