package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:secret@tcp(localhost:3306)/gorm")
	if err != nil {
		panic(err)
	}

	// query(db)
	typeSafeMapping(db)
}
