package main

import (
	"database/sql"
	"fmt"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/factories"
	negativations "github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/services/negativation"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
)

// TODO:
// - improve error handling and messaging returns
// 		- create a library for:
//			- handling errors
// 			- returning messages for either the user &| internally
// - correct mock end-point tests
// 		- by some reason they failing internally at lines:
// 			- 69, 94, 119, 154 and 191
// - implement authentication middleware
// - implement .env on legacy api controller client-server
// - swagger variables and names should be created dynamically with .env variables
// - improve swagger parameter description
// - fix application dockerization

func main() {
	fmt.Println("HERE ->", os.Getenv("DB_DRIVER"))
	config, err := util.Config.LoadConfig(".")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("HERE ->", config.DBDriver, config.DBDataSourceName)

	dbConn, err := sql.Open(config.DBDriver, config.DBDataSourceName)
	if err != nil {
		panic(err)
	}

	swagger := factories.Swagger.Generate()
	go func(eng *gin.Engine) {
		// TODO: .env variable for this path address
		eng.Run("localhost:8010")
	}(swagger)

	legacyProxy := factories.LegacyProxy.NewLegacyServer()
	go func(legacyProxy *factories.SLegacyServer, address string) {
		err = legacyProxy.Start(address)
		if err != nil {
			panic(err)
		}
	}(legacyProxy, config.LegacyServerAddress)

	negStore := negativations.NewStore(dbConn)
	negServer := factories.Negativations.NewServer(negStore)
	err = negServer.Start(config.NegativationsServerAddress)
	if err != nil {
		panic(err)
	}
}
