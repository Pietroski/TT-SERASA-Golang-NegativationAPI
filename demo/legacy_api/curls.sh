#!/usr/bin/env bash

curl --request GET \
  --url http://localhost:8009/v1-legacy/ping

curl --request GET \
  --url http://localhost:8009/v1-legacy/negativated/1

curl --request GET \
  --url 'http://localhost:8009/v1-legacy/list-negativated?page_number=1&page_size=5' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data =
