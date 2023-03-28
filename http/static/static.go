package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.HandleFunc("/", fs.ServeHTTP)

	log.Println("Executing...")
	log.Fatalln(http.ListenAndServe(":3001", nil))
}
