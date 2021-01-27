package handler

import (
	"context"
	"github.com/sliaptsou/backend/internal/adapter"
	"github.com/sliaptsou/backend/internal/entity"
	"github.com/sliaptsou/backend/internal/repo"
	"github.com/sliaptsou/backend/internal/service"
	"github.com/sliaptsou/backend/proto"
)

type Backend struct{}

var _ proto.BackendServer = (*Backend)(nil)

func NewBackendServer() *Backend {
	return &Backend{}
}

func (t *Backend) GetOne(ctx context.Context, request *proto.GetOneItemRequest) (*proto.GetOneItemResponse, error) {
	rsp := &proto.GetOneItemResponse{
		Id: request.Id,
	}

	db, err := repo.GetDb()
	if err != nil {
		return nil, err
	}

	e, err := service.NewEntityService(ctx, db).GetByID(request.Id)
	if err != nil {
		return nil, err
	}

	err = adapter.NewEntityGrpcAdapter().ToGrpc(e, rsp)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *Backend) Create(ctx context.Context, request *proto.CreateRequest) (*proto.GetOneItemResponse, error) {
	db, err := repo.GetDb()
	if err != nil {
		return nil, err
	}

	e := &entity.Entity{Name: request.Name}
	if err = service.NewEntityService(ctx, db).Create(e); err != nil {
		return nil, err
	}

	return &proto.GetOneItemResponse{Name: e.Name, Id: e.ID}, nil
}
