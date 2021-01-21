.PHONY: pull build push run_backend network clear proto vet test lint count check

PATH_TO_LINTER=$(go env GOPATH)/bin
NETWORK=example
REGISTRY=sliaptsou
GATEWAY_NAME=gateway
BACKEND_NAME=backend
GATEWAY_IMAGE=$(REGISTRY)/gateway
BACKEND_IMAGE=$(REGISTRY)/backend
API_PORT=8080
SVC_HOST=$(BACKEND_NAME)
SVC_PORT=8080
LOCAL_PORT=8080

pull:
	@docker pull $(GATEWAY_IMAGE):latest
	@docker pull $(BACKEND_IMAGE):latest

build:
	docker build -t $(BACKEND_IMAGE):latest . && \
	docker build -t $(GATEWAY_IMAGE):latest ./gateway

push:
	docker push $(BACKEND_IMAGE):latest
	docker push $(GATEWAY_IMAGE):latest

run: run_backend run_gateway
	@echo "Application run on http://localhost:$(API_PORT)"

run_backend:
	@docker run -d \
               --name $(BACKEND_NAME) \
               --env SVC_HOST=$(SVC_HOST) \
               --env SVC_PORT=$(SVC_PORT) \
               --network $(NETWORK) \
               $(BACKEND_IMAGE):latest

run_gateway: run_backend
	@docker run -d \
               --name $(GATEWAY_NAME) \
               --publish $(LOCAL_PORT):$(API_PORT) \
               --env API_PORT=$(API_PORT) \
               --env SVC_HOST=$(SVC_HOST) \
               --env SVC_PORT=$(SVC_PORT) \
               --network $(NETWORK) \
               $(GATEWAY_IMAGE):latest

check:
	@curl http://localhost:$(API_PORT)/health

count:
	@curl http://localhost:$(API_PORT)/count

network:
	@docker network create $(NETWORK)

clear:
	@docker stop $(BACKEND_NAME) $(GATEWAY_NAME)
	@docker rm $(BACKEND_NAME) $(GATEWAY_NAME)

proto:
	protoc --go_out=plugins=grpc:. ./proto/backend.proto && \
	protoc-go-inject-tag -input=proto/backend.pb.go

vet:
	go vet -composites=false ./...

fmt: vet
	gofmt -s -w .

test:
	go test ./... -v --cover

lint:
	$(PATH_TO_LINTER)/golangci-lint run -v --config golangci.yml
