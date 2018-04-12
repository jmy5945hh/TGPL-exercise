/*
* Exercise 1.9: Modify fetch to also print the HTTP status code,found in resp.Status.
 */

package main

import (
	"fmt"
	"net/http"
	"os"
)

func fetchWithStatusCode() {
	url := "https://jmy5945hh.github.io/"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(resp.Status)

	//...
}
