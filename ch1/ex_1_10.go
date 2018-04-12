/*
* Exercise 1.10: Find a website that produces a large amount of data.
* Investigate caching by running fetchall twice in succession to see
* whether the reported time changes much. Do you get the same content each time?
* Modify fetchall to print its output to a file so it can be examined.
 */

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func fetchAll() {
	start := time.Now()
	ch := make(chan string)

	urls := []string{"http://fund.eastmoney.com/", "http://fund.eastmoney.com/"}

	for _, url := range urls {
		go fetchConCurr(url, ch) // start a goroutine
	}

	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchConCurr(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
