package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("== Relacionais ==")
	fmt.Println("Strings: ", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	date1 := time.Unix(0, 0)
	date2 := time.Unix(0, 0)
	fmt.Println("Datas: ", date1 == date2)
	fmt.Println("Datas: ", date1.Equal(date2))

	type Pessoa struct {
		Nome  string
		Idade int
	}

	pessoa1 := Pessoa{"João", 20}
	pessoa2 := Pessoa{"João", 20}

	fmt.Println("Struct: ", pessoa1 == pessoa2)
}
