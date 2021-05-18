GO=go
DOCKER=docker
DOCKER_COMPOSE=docker-compose
MIGRATE=migrate
SQLC=sqlc
SWAG=swag

docker-container:
	$(DOCKER_COMPOSE) up -d

stop-docker-container:
	$(DOCKER_COMPOSE) down

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

all: docker-container migrations test swagger run

 .PHONY: docker-container stop-docker-container migrations reverse-migrations sqlc mock test-database test swagger run
