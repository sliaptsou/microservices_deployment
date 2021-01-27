package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sliaptsou/backend/internal/entity"
	"github.com/sliaptsou/backend/internal/repo/entityRepo"
)

type EntityService struct {
	ctx              context.Context
	EntityRepository entityRepo.Repository
}

func NewEntityService(ctx context.Context, db *sqlx.DB) EntityService {
	return EntityService{
		ctx:              ctx,
		EntityRepository: entityRepo.New(ctx, db),
	}
}

func (s EntityService) GetByID(id int32) (*entity.Entity, error) {
	entity, err := s.EntityRepository.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Not found")
		}
		return nil, err
	}
	return entity, err
}

func (s EntityService) Create(entity *entity.Entity) error {
	return s.EntityRepository.Create(entity)
}
