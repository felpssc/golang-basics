package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@/golangproject")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO users(name) values(?)")

	stmt.Exec("Felipe")
	stmt.Exec("Diego")

	res, err := stmt.Exec("Neymar")

	if err != nil {
		panic(err)
	}

	fmt.Println(res.LastInsertId())
}
