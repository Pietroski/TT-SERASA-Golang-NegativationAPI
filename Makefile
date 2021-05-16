GO=go
DOCKER=docker
DOCKER_COMPOSE=docker-compose
MIGRATE=migrate
SQLC=sqlc

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

test:
	go test -v -cover ./...

all: docker-container migrations test

 .PHONY: docker-container migrations test
