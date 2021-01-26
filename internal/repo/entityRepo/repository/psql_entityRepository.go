package repository

import (
	"context"
	"log"

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

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	q := psql.Select("*").From("entity").Where(sq.Eq{"id": id})
	query, args, err := q.ToSql()
	log.Printf("ToSql: %+v", err)
	if err != nil {
		return nil, err
	}
	log.Printf("query: %+v", query)
	log.Printf("args: %+v", args)

	err = repo.pool.GetContext(repo.ctx, &m, query, args...)

	log.Printf("GetContext: %+v", err)

	m.LoadToEntity(&e)

	return &e, err
}

func (repo *psqlWarehouseRepository) Create(e *entity.Entity) error {
	m := model.Entity{}
	m.LoadFromEntity(*e)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query, args, err := psql.Insert("entity").
		Columns("name").
		Values(m.Name).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return err
	}

	rows := repo.pool.QueryRowxContext(repo.ctx, query, args...)
	err = rows.StructScan(&m)
	if err != nil {
		return err
	}

	m.LoadToEntity(e)

	return nil
}
