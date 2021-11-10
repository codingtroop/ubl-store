package interfaces

import (
	"context"

	"github.com/codingtroop/ubl-store/pkg/entities"
)

type UblRepositoryInterface interface {
	Get(cntxt context.Context, id string) (*entities.UblEntity, error)
	Insert(cntxt context.Context, id string, ubl entities.UblEntity) error
	Delete(cntxt context.Context, id string) error
}
