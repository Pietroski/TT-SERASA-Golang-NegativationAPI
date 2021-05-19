#!/usr/bin/env bash

set -e

sleep 10

echo "run db migration"
/app/migrate -path /app/migrations -database "postgresql://serasa:serasa_psql@localhost:5432/tt_serasa?sslmode=disable" -verbose up

echo "start the app"
exec "$@"
