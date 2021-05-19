package negativations

import (
	"database/sql"
	"fmt"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/util"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
)

func TestMain(m *testing.M) {
	fmt.Println("Testing Database Connection on -> TestMain")

	config, err := util.Config.LoadConfig("../../../")
	if err != nil {
		fmt.Println(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBDataSourceName)
	if err != nil {
		panic(err)
	}

	testQueries = New(conn)

	fmt.Println()
	os.Exit(m.Run())
}
