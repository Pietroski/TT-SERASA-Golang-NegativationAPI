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

### Viper commands

### Gin-Gonic commands

### go-mock commands

- After installing and configuring it, to create a database mock for your project:
```shell
mockgen -package mock_negativation -destination internal/services/negativation/mock/store.go github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/services/negativation Store
```

### Using Swaggo

- Installing Swaggp
```shell
go get -u github.com/swaggo/swag/cmd/swag
```

- After installing run:
```shell
swag init
```

- if your main.go is not on your root directory as this project run:
```shell
swag init -g path/to/your/main.go
```

-   - in our case:
    ```shell
    swag init -g cmd/main.go
    ```

- After configuring sawggo and whilst running the server, to go look at the swagger files go to:
- http://localhost:8010/swagger/index.html

