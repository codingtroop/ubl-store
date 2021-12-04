package helpers

import (
	"github.com/codingtroop/ubl-store/pkg/config"
	"github.com/codingtroop/ubl-store/pkg/helpers/interfaces"
)

var (
	ublFolder    = "ubl"
	attachFolder = "attachment"
)

func ResolveStorage(sc config.StorageConfiguration) (ubl interfaces.Storer, attach interfaces.Storer) {

	if sc.Filesystem.DataPath != "" {
		ubl = NewIOStorer(sc.Filesystem.DataPath + "/" + ublFolder)
		attach = NewIOStorer(sc.Filesystem.DataPath + "/" + attachFolder)
		return
	}

	if sc.S3.Bucket != "" {
		ubl = NewS3Storer(sc.S3.Endpoint, ublFolder, sc.S3.Bucket)
		attach = NewS3Storer(sc.S3.Endpoint, attachFolder, sc.S3.Bucket)
		return
	}

	return nil, nil
}
