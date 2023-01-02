package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func titles(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)

			regex, _ := regexp.Compile(`<title>(.*?)</title>`)

			title := regex.FindStringSubmatch(string(html))[1]

			c <- title
		}(url)
	}

	return c
}

func main() {
	t1 := titles("https://www.google.com", "https://www.yahoo.com")
	t2 := titles("https://www.youtube.com", "https://www.globo.com")

	fmt.Println("First: ", <-t1, "|", <-t2)
	fmt.Println("Second: ", <-t1, "|", <-t2)
}
