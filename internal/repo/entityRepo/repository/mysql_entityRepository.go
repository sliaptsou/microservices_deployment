package repository

import (
	"context"
	"github.com/sliaptsou/backend/internal/entity"
	"github.com/sliaptsou/backend/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type psqlWarehouseRepository struct {
	ctx  context.Context
	pool *sqlx.DB
}

func NewPsqlWarehouseRepository(ctx context.Context, pool *sqlx.DB) *psqlWarehouseRepository {
	return &psqlWarehouseRepository{
		ctx:  ctx,
		pool: pool,
	}
}

func (repo *psqlWarehouseRepository) GetByID(id int32) (*entity.Entity, error) {
	m := model.Entity{}
	e := entity.Entity{}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Question)
	q := psql.Select("*").From("entity").Where(sq.Eq{"id": id})
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	err = repo.pool.GetContext(repo.ctx, &m, query, args...)

	m.LoadToEntity(&e)

	return &e, err
}

func (repo *psqlWarehouseRepository) Create(e *entity.Entity) error {
	m := model.Entity{}
	m.LoadFromEntity(*e)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Question)
	query, args, err := psql.Insert("entity").
		Columns("name").
		Values(m.Name).
		//Suffix("RETURN *").
		ToSql()

	if err != nil {
		return err
	}

	rows := repo.pool.QueryRowxContext(repo.ctx, query, args...)
	err = rows.Err()
	if err != nil {
		return err
	}

	//m.LoadToEntity(e)

	return nil
}
