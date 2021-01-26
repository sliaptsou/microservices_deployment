module github.com/sliaptsou/backend

go 1.15

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/favadi/protoc-go-inject-tag v1.1.0 // indirect
	github.com/golang-migrate/migrate v3.5.4+incompatible // indirect
	github.com/jmoiron/sqlx v1.3.1
	github.com/lib/pq v1.3.0
	github.com/micro/go-micro/v2 v2.9.1 // indirect
	github.com/sliaptsou/backend/proto v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210122163508-8081c04a3579 // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
)

replace github.com/sliaptsou/backend/proto => ./proto
