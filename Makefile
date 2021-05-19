GO=go
DOCKER=docker
DOCKER_COMPOSE=docker-compose
MIGRATE=migrate
SQLC=sqlc
SWAG=swag

SUDO=sudo

docker-build:
	docker build -t pietroski/tt_serasa_golang_server .

build-docker-containers:
	$(DOCKER_COMPOSE) up -d --build

build-docker-containers-with-logs:
	$(DOCKER_COMPOSE) up --build

docker-containers:
	$(DOCKER_COMPOSE) up -d

docker-containers-with-logs:
	$(DOCKER_COMPOSE) up

stop-docker-containers:
	$(DOCKER_COMPOSE) down

remove-docker-containers:
	$(DOCKER) rm tt_serasa_golang_server tt_serasa_postgres

clean-local-db:
	$(SUDO) rm -rf .db

migrations:
	$(MIGRATE) -path internal/datastore/postgreSQL/migrations -database "postgresql://serasa:serasa_psql@localhost:5432/tt_serasa?sslmode=disable" -verbose up

reverse-migrations:
	$(MIGRATE) -path internal/datastore/postgreSQL/migrations -database "postgresql://serasa:serasa_psql@localhost:5432/tt_serasa?sslmode=disable" -verbose down

sqlc:
	$(SQLC) generate

mock:
	mockgen -package mock_negativation -destination internal/services/negativation/mock/store.go github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/services/negativation Store

swagger:
	$(SWAG) init -g cmd/main.go

test-database:
	$(GO) test -v -cover ./internal/services/negativation/

test:
	$(GO) test -v -cover ./...

run:
	$(GO) run cmd/main.go

all: docker-containers migrations test swagger run

.PHONY: docker-build build-docker-containers build-docker-containers-with-logs docker-containers docker-containers-with-logs stop-docker-containers remove-docker-containers clean-local-db migrations reverse-migrations sqlc mock test-database test swagger run
