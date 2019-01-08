package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(uri string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(uri)

	if err != nil {
		ch <- fmt.Sprint(err)

		return
	}

	file, err := os.Create(url.QueryEscape(uri))

	if err != nil {
		ch <- fmt.Sprintf("while creating file: %v", err)

		return
	}

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while rading %s: %v", uri, err)

		return
	}

	if err := file.Close(); err != nil {
		ch <- fmt.Sprintf("while closing file: %v", err)

		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, uri)
}
