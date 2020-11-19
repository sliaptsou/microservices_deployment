package handler

import (
	"context"

	"github.com/sliaptsou/backend/proto"
)

type Backend struct{}

var _ proto.BackendServer = (*Backend)(nil)

func NewBackendServer() *Backend {
	return &Backend{}
}

func (t *Backend) HealthCheck(_ context.Context, _ *proto.Empty) (rsp *proto.HttpCode, err error) {
	rsp = &proto.HttpCode{Code: 200}
	return rsp, nil
}