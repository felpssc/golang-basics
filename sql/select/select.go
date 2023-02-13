package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:admin@/golangproject")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query("SELECT * FROM users WHERE id > ?", 0)
	defer rows.Close()

	var users []user

	for rows.Next() {
		var u user

		rows.Scan(&u.id, &u.name)

		users = append(users, u)
	}

	fmt.Println(users)
}
