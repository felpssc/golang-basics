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

	stmt, _ := db.Prepare("UPDATE users SET name = ? WHERE id = ?")

	stmt.Exec("Cleiton", 1)
	stmt.Exec("Clovis", 10)

	stmt, _ = db.Prepare("DELETE FROM users WHERE id = ?")

	stmt.Exec(9)
	stmt.Exec(1)
}
