package interfaces

import (
	"context"

	"github.com/codingtroop/ubl-store/pkg/entities"
	"github.com/google/uuid"
)

type UblRepository interface {
	Get(cntxt context.Context, uuid uuid.UUID) (*entities.UblEntity, error)
	Insert(cntxt context.Context, ubl entities.UblEntity) error
	Delete(cntxt context.Context, uuid uuid.UUID) error
}
