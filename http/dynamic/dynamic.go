package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeNow(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")

	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/timeNow", timeNow)
	log.Fatal(http.ListenAndServe(":3002", nil))
}
