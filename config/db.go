package config

import (
	"database/sql"
	"fmt"
)

func ConnectDb() *sql.DB {
	dsn := "user:password@tcp(localhost:3319)/ae"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success connection to MySql")

	return db
}
