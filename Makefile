PATH_TO_LINTER=$(go env GOPATH)/bin
NETWORK=example
GATEWAY_NAME=gateway
BACKEND_NAME=backend
API_PORT=8090
SVC_HOST=$(BACKEND_NAME)
SVC_PORT=8091

build:
	docker build -t $(BACKEND_NAME):latest . && \
	docker build -t $(GATEWAY_NAME):latest ./gateway

run:
	docker network create $(NETWORK) && \
    docker run -d \
               --name $(BACKEND_NAME) \
               --env SVC_HOST=$(SVC_HOST) \
               --env SVC_PORT=$(SVC_PORT) \
               --network $(NETWORK) \
               $(BACKEND_NAME):latest &&\
    docker run -d \
               --name $(GATEWAY_NAME) \
               --publish 8090:$(API_PORT) \
               --env API_PORT=$(API_PORT) \
               --env SVC_HOST=$(SVC_HOST) \
               --env SVC_PORT=$(SVC_PORT) \
               --network $(NETWORK) $(GATEWAY_NAME):latest

clear:
	docker stop $(BACKEND_NAME) $(GATEWAY_NAME) && \
	docker rm $(BACKEND_NAME) $(GATEWAY_NAME) && \
    docker network rm $(NETWORK)

.PHONY: proto
proto:
	protoc --go_out=plugins=grpc:. ./proto/backend.proto && protoc-go-inject-tag -input=proto/backend.pb.go

.PHONY: vet
vet:
	go vet -composites=false ./...

fmt: vet
	gofmt -s -w .

.PHONY: test
test:
	go test ./... -v --cover

.PHONY: lint
lint:
	$(PATH_TO_LINTER)/golangci-lint run -v --config golangci.yml