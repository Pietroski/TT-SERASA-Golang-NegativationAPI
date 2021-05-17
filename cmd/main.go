package main

import (
	"database/sql"
	"fmt"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/factories"
	negativations "github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/services/negativation"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/util"
	_ "github.com/lib/pq"
)

// TODO:
// - improve error handling and messaging returns
// 		- create a library for:
//			- handling errors
// 			- returning messages for either the user &| internally
// -

func main() {
	config, err := util.Config.LoadConfig(".")
	if err != nil {
		fmt.Println(err)
	}

	dbConn, err := sql.Open(config.DBDriver, config.DBDataSourceName)
	if err != nil {
		panic(err)
	}

	negStore := negativations.NewStore(dbConn)
	negServer := factories.Negativations.NewServer(negStore)

	err = negServer.Start(config.NeggativationsServerAddress)
	if err != nil {
		panic(err)
	}
}
