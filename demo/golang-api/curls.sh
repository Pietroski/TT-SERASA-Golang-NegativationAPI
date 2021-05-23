#!/usr/bin/env bash

curl --request GET \
  --url http://localhost:8008/v2/ping

curl --request POST \
  --url http://localhost:8008/v2/negativate \
  --header 'Content-Type: application/json' \
  --data '{
	"companyDocument": "59291534043167",
	"companyName": "ABCJ S.A.",
	"customerDocument": "51537477567",
	"value": 5432.23,
	"contract": "bc063153-asfd-fsa5-v4cx-0d069a42065b",
	"debtDate": "2015-12-13T20:32:51-01:00",
	"inclusionDate": "2019-11-13T20:32:51-02:00"
}'

curl --request PUT \
  --url http://localhost:8008/v2/negativated \
  --header 'Content-Type: application/json' \
  --data '{
	"id": 139,
	"companyDocument": "59291534043165",
	"companyName": "ACD S.A.",
	"customerDocument": "51537477567",
	"value": 5432.20,
	"contract": "bc063153-asfd-fsa5-v4cx-0d069a42065b",
	"debtDate": "2015-12-13T20:32:51-01:00",
	"inclusionDate": "2019-11-13T20:32:51-02:00"
}'

curl --request GET \
  --url 'http://localhost:8008/v2/negativated?page_number=1&page_size=5' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data =

curl --request GET \
  --url http://localhost:8008/v2/negativated/139

curl --request DELETE \
  --url http://localhost:8008/v2/negativated/139
