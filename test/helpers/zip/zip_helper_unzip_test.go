package test

import (
	"context"
	"testing"

	"github.com/codingtroop/ubl-store/pkg/helpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Zip_UnZip_Ok(t *testing.T) {
	zh := helpers.NewZipper()

	fileName := uuid.New().String()
	data := []byte(uuid.New().String())

	compressed, err := zh.Compress(context.TODO(), fileName, data)

	if err != nil {
		t.Fatal(err)
	}

	decompressed, err := zh.Decompress(context.TODO(), compressed)

	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, data, decompressed, "they should be equal")
}
