package main

import "fmt"

func typeof(arg interface{}) string {
	switch arg.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float32, float64:
		return "float"
	default:
		return "nil"
	}
}

func main() {
	fmt.Println(typeof("oi"))
}
