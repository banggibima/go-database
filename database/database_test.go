package database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func TestGetConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
}
