package model

import (
	"github.com/sliaptsou/backend/internal/entity"
)

type Entity struct {
	ID   int32  `db:"id"`
	Name string `db:"name"`
}

func (w Entity) LoadToEntity(e *entity.Entity) {
	e.ID = w.ID
	e.Name = w.Name
}

func (w *Entity) LoadFromEntity(e entity.Entity) {
	w.ID = e.ID
	w.Name = e.Name
}
