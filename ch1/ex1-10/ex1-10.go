// Investigate caching by running fetchall twice in succession. Modify fetchall
// to print its output to a file so it can be examined.
//
// http://infowars.com
// 1st try: 1.5s
// 2nd try: 1.53s
//
// Solved by Tadeusz Tomoszek 13.7.2018

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	f, err := os.Create("fetchall")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
		return
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch, f) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, f *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <-fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

