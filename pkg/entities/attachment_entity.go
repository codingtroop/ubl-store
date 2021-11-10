package entities

import (
	"time"

	"github.com/google/uuid"
)

type AttachmentEntity struct {
	ID      uuid.UUID
	Created time.Time
	UblID   uuid.UUID
	Hash    string
}
