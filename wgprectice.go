package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"http://google.com",
	"http://twitter.com",
	"http://yandex.com",
}

func main() {
	fmt.Println("Practice started")
	http.HandleFunc("/", getStatusofURL)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getStatusofURL(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%+v\n", resp)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
