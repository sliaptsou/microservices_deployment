module github.com/sliaptsou/backend

go 1.14

require (
	github.com/sliaptsou/backend/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.33.2
)

replace github.com/sliaptsou/backend/proto => ./proto
