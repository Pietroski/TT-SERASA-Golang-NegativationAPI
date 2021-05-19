# Serasa's Technical Test

## Considerações finais

- No início do arquivo cmd/main.go tem alguns TODOS que eu gostaria ter implementado mas que não consegui por falta de tempo.
- Apesar de relativamente simples esse test for relativamente longo e apresenta bastante detalhezinhos.
- Espero que pelo menos sirva como amostra.

- Apesar de eu não ter conseguido fazer funcionar a aplicação no docker junto ao container do postgres os arquivos ficarão aí para mostrar a tentativa.
- Por ter gastado um bom tempo tentando corrigir a imagem do docker, não consegui terminar de introduzir os inputs dos endpoints no Swagger e por isso fiz uns aquivos de consulta em ./demo.

- Para rodar a aplicação com o banco de dados já "migracionado", basta apenas rodar
```shell
make all
```

- Para rodar algum comando específico basta apenas consultar o arquivo Makefile.

- Senti que teve bastante tasks de devops nesse teste técnico. 
- Fiz uma boa bateria de testes unitários porém não consegui terminar de corrigir os testes de mock.
- Os testes de banco de dados estão passando inclusive nas pipelines do github actions.

- Para rodar a aplicação legada basta apenas seguir a documentação do repositório dela.
  https://github.com/Pietroski/TT-SERASA-NodeJS-LegacyNegativationAPI
  
- Eu sei que no teste pedia para que não fosse mergeado o request porém eu decidi fazê-lo para garantir que eu não perdesse algumas coisas.
- Inclusive, para que vocês pudessem ler esse README.md
- Eu sei que tem milhares de outras formas diferentes e mais fácil de se executar isso, mas decidi que no final fosse tudo "mergeado".

- Enfim, ficaram alguns detalhezinhos mas como o tempo está curto no momento foi o que eu consegui entregar em muito pouco tempo de mão na massa.
- Espero que gostem mas vou entender também caso não gostem.
- O prazo era de 7 dias e foi isso o que eu consegui fazer; começei a montagem dessa API domingo a noite em decorrência de outros testes e atividades.
- Pude trabalhar nela apenas no Domingo(como já comentado), segunda e terça a noite. No total foram aproximadamente 8-12 horas de trabalho.

- Sobre as escolhas das libs, obviamente foram feitas para tentar manter o zen do Go; descomplicada compilação, descomplicada execução e programação. 
- A maioria das bibliotecas escolhidas foi escolhida por manter performance e type-checking. Para mais informações basta apenas me perguntar.

- A arquitetura que eu tento aplicar em meus projeto de go é baseada nesse cara -> https://github.com/golang-standards/project-layout e internamente sempre tempo ainda seguir os princípios de SOLID e Clean Architecture, vide ./test/mock(além de todo o restante do projeto) por exemplo.

- Se eu não esqueci de nenhum detalhe basta só me perguntar, vou adorar responder, mesmo que eu não saiba de algo. 
- Caso eu não saiba, vou procurar ir atrás da infomrção; de qualquer modo vai valer o aprendizado.

- Grato e tenham todos uma ótima semana!!

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
