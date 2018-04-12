/*
* Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes to dst.
* Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring
* a buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
 */
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func fetch() {
	url := "https://jmy5945hh.github.io/"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
}