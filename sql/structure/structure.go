package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	db, err := sql.Open("mysql", "root:admin@/")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	database := "golangproject"
	table := "users"

	createDatabaseQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", database)
	useDatabaseQuery := fmt.Sprintf("USE %s", database)
	dropTableQuery := fmt.Sprintf("DROP TABLE IF EXISTS %s", table)
	createUsersTableQuery := fmt.Sprintf(`CREATE TABLE %s (
		id integer auto_increment,
		name varchar(100),
		PRIMARY KEY (id)
	)`, table)

	exec(db, createDatabaseQuery)
	exec(db, useDatabaseQuery)
	exec(db, dropTableQuery)
	exec(db, createUsersTableQuery)
}
