package interfaces

import (
	"context"

	"github.com/codingtroop/ubl-store/pkg/entities"
	"github.com/google/uuid"
)

type AttachmentRepository interface {
	Get(cntxt context.Context, uuid uuid.UUID) (*entities.AttachmentEntity, error)
	Insert(cntxt context.Context, attachment entities.AttachmentEntity) error
	Delete(cntxt context.Context, uuid uuid.UUID) error
}
