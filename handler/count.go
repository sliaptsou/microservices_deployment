package handler

import (
	"context"
	"sync/atomic"

	"github.com/sliaptsou/backend/proto"
)

var count int64

func (t *Backend) GetQueryCount(_ context.Context, _ *proto.Empty) (*proto.CountResponse, error) {
	atomic.AddInt64(&count, 1)
	return &proto.CountResponse{Id: count}, nil
}
