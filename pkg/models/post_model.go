package models

import "github.com/google/uuid"

type PostModel struct {
	ID   uuid.UUID
	Data []byte
}
