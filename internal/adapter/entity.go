package adapter

import (
	"github.com/sliaptsou/backend/internal/entity"
	"github.com/sliaptsou/backend/proto"
)

type EntityGrpcAdapter struct {
}

func NewEntityGrpcAdapter() *EntityGrpcAdapter {
	return &EntityGrpcAdapter{}
}

func (p EntityGrpcAdapter) ToGrpc(r *entity.Entity, rsp *proto.GetOneItemResponse) error {
	rsp.Id = r.ID
	rsp.Name = r.Name
	return nil
}
