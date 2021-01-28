.PHONY: pull build push run_backend network clear proto vet test lint count check

PATH_TO_LINTER=$(shell go env GOPATH)/bin
NETWORK=example
REGISTRY=sliaptsou
GATEWAY_NAME=gateway
BACKEND_NAME=backend
DATABASE_NAME=database
GATEWAY_IMAGE=$(REGISTRY)/gateway
BACKEND_IMAGE=$(REGISTRY)/backend
#DATABASE_IMAGE=postgres:12.5
DATABASE_IMAGE=mysql:5.7
API_PORT=8080
SVC_HOST=$(BACKEND_NAME)
SVC_PORT=8080
LOCAL_PORT=8080
NAMESPACE=example-ns

POSTGRES_USER=appuser
POSTGRES_PASSWORD=application_user_pass
POSTGRES_DB=application_db
POSTGRES_SSL_MODE=disable
#POSTGRES_PORT=5432
POSTGRES_PORT=3306
TAG=0.2

pull:
	@docker pull $(GATEWAY_IMAGE):$(TAG)
	@docker pull $(BACKEND_IMAGE):$(TAG)
	@docker pull $(DATABASE_IMAGE)

build: build_backend build_gateway
	@echo "Builded"

build_backend:
	docker build --no-cache \
	 -t $(BACKEND_IMAGE):latest \
	 -t $(BACKEND_IMAGE):$(TAG) \
	  .

build_gateway:
	docker build --no-cache \
	 -t $(GATEWAY_IMAGE):latest \
	 -t $(GATEWAY_IMAGE):$(TAG) \
	 ./gateway

push:
	docker push --all-tags $(BACKEND_IMAGE)
	docker push --all-tags $(GATEWAY_IMAGE)

run: run_backend run_gateway
	@echo "Application run on http://localhost:$(API_PORT)"


#run_db:
#	@docker run -d \
#               --name $(DATABASE_NAME) \
#               --env POSTGRES_USER=$(POSTGRES_USER) \
#               --env POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
#               --env POSTGRES_DB=$(POSTGRES_DB) \
#               -p 5433:$(POSTGRES_PORT) \
#               --network $(NETWORK) \
#               $(DATABASE_IMAGE)

run_db:
	@docker run -d \
               --name $(DATABASE_NAME) \
               --env MYSQL_USER=$(POSTGRES_USER) \
               --env MYSQL_PASSWORD=$(POSTGRES_PASSWORD) \
               --env MYSQL_DATABASE=$(POSTGRES_DB) \
               --env MYSQL_ROOT_PASSWORD=root \
               -p 3306:$(POSTGRES_PORT) \
               --network $(NETWORK) \
               $(DATABASE_IMAGE)

run_backend:
	@docker run -d \
               --name $(BACKEND_NAME) \
               --env SVC_HOST=$(SVC_HOST) \
               --env SVC_PORT=$(SVC_PORT) \
               --env DB_HOST=$(DATABASE_NAME) \
               --env DB_PORT=$(POSTGRES_PORT) \
               --env DB_USER=$(POSTGRES_USER) \
               --env DB_USER_PASS=$(POSTGRES_PASSWORD) \
               --env DB_NAME=$(POSTGRES_DB) \
               --env DB_SSL_MODE=$(POSTGRES_SSL_MODE) \
               --env DB_MAX_OPEN_CONN=2 \
               --network $(NETWORK) \
               $(BACKEND_IMAGE):latest

run_gateway:
	@docker run -d \
               --name $(GATEWAY_NAME) \
               --publish $(LOCAL_PORT):$(API_PORT) \
               --env API_PORT=$(API_PORT) \
               --env SVC_HOST=$(SVC_HOST) \
               --env SVC_PORT=$(SVC_PORT) \
               --network $(NETWORK) \
               $(GATEWAY_IMAGE):latest

run_migration:
	migrate -path internal/migration -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5433/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSL_MODE) up

create_migration:
	migrate create -ext sql -dir internal/migration -seq create_entity_table

network:
	@docker network create $(NETWORK)

clear:
	@docker stop $(BACKEND_NAME) $(GATEWAY_NAME)
	@docker rm $(BACKEND_NAME) $(GATEWAY_NAME)

check:
	@curl http://localhost:$(API_PORT)/healthz

count:
	@curl http://localhost:$(API_PORT)/count




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



#minikube
.PHONY: apply delete list url health-check
apply:
	kubectl apply -f deployments/manifests --recursive=true --namespace=$(NAMESPACE)

delete:
	kubectl delete -f deployments/manifests --recursive=true --namespace=$(NAMESPACE)

list:
	kubectl get pods --namespace=$(NAMESPACE)

url:
	minikube service gateway-example-chart --url --namespace=$(NAMESPACE)


#	@curl $(shell minikube service gateway --url --namespace=$(NAMESPACE))/healthz
health-check:
	@curl $(shell minikube service gateway-example-chart --url --namespace=$(NAMESPACE))/health

namespace:
	kubectl create namespace $(NAMESPACE)

#HELM
.PHONY: helm-install helm-delete helm-upgrade describe
helm-install:
	helm install example-chart deployments/example-chart --namespace=$(NAMESPACE)

helm-delete:
	helm delete example-chart --namespace=$(NAMESPACE)

helm-upgrade:
	helm upgrade example-chart deployments/example-chart -f deployments/example-chart/values.yaml --namespace=$(NAMESPACE)

describe:
	kubectl describe pods $(POD) --namespace=$(NAMESPACE)