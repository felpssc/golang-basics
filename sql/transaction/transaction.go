package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@/golangproject")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	transaction, _ := db.Begin()

	stmt, _ := transaction.Prepare("INSERT INTO users(id, name) VALUES(?,?)")

	stmt.Exec(9, "Carlos")
	stmt.Exec(10, "Pedro")

	_, err = stmt.Exec(1, "Lucas")

	if err != nil {
		transaction.Rollback()
		panic(err)
	}

	transaction.Commit()
}
