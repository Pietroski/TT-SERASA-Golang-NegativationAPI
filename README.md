# Serasa's Technical Test

## general list of commands

### Migrations

- for consulting migrate help command:
```shell
migrate -help
```

- for initiating/creating migration files:
```shell
migrate create -ext sql -dir path/to/your/migration/files -seq init_schema
```
-
    - in our case:
    ```shell
    migrate create -ext sql -dir internal/datastore/postgreSQL/migrations/ -seq init_schema
    ```
  
- for migrating up verbosely:
```shell
migrate -path path/to/your/db/migration/file -database "your_database_url" -verbose up
```

-   -for this test example:
    ```shell
    migrate -path internal/datastore/postgreSQL/migrations/ -database "postgresql://serasa:serasa_psql@localhost:5432/tt_serasa?sslmode=disable" -verbose up
    ```

- for migrating down:
```shell
migrate -path path/to/your/db/migration/file -database "your_database_url" -verbose down
```

-   -for this test example:
    ```shell
    migrate -path internal/datastore/postgreSQL/migrations/ -database "postgresql://serasa:serasa_psql@localhost:5432/tt_serasa?sslmode=disable" -verbose down
    ```
    
### SQLC commands

- To start a sqlc configuration file on your root directory:
```shell
sqlc init
```

- To generate the go files after properly configuring and writing the desired queries:
```shell
sqlc generate
```