package helpers

import (
	"github.com/codingtroop/ubl-store/pkg/config"
	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

func ResolveStorage(sc config.StorageConfiguration) (ubl interfaces.Storer, attach interfaces.Storer) {

	if sc.Filesystem.AttachmentPath != "" && sc.Filesystem.UblPath != "" {
		ubl = NewIOStorer(sc.Filesystem.UblPath)
		attach = NewIOStorer(sc.Filesystem.AttachmentPath)
		return
	}

	if sc.S3.AttachmentPath != "" && sc.S3.Bucket != "" && sc.S3.Endpoint != "" && sc.S3.UblPath != "" {
		ubl = NewS3Storer(sc.S3.Endpoint, sc.S3.UblPath, sc.S3.Bucket)
		attach = NewS3Storer(sc.S3.Endpoint, sc.S3.AttachmentPath, sc.S3.Bucket)
		return

	}

	return nil, nil
}
