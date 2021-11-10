package entities

import (
	"time"

	"github.com/google/uuid"
)

type UblEntity struct {
	ID      uuid.UUID
	Created time.Time
}