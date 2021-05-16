package negativations

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbDataSourceName = "postgresql://serasa:serasa_psql@localhost:5432/tt_serasa?sslmode=disable"
)

var (
	testQueries *Queries
)

func TestMain(m *testing.M) {
	fmt.Println("Testing Database Connection on -> TestMain")

	conn, err := sql.Open(dbDriver, dbDataSourceName)
	if err != nil {
		panic(err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
