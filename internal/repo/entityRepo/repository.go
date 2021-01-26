package entityRepo

import (
	"context"

	"github.com/sliaptsou/backend/internal/entity"
	"github.com/sliaptsou/backend/internal/repo/entityRepo/repository"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetByID(id int32) (*entity.Entity, error)
	Create(*entity.Entity) error
}

func New(ctx context.Context, pool *sqlx.DB) Repository {
	return repository.NewPsqlWarehouseRepository(ctx, pool)
}
