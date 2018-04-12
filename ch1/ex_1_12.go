/*
* Exercise 1.12: Modify the Lissajous server to read parameter values from the URL.
* For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20
* sets the number of cycles to 20 instead of the default 5.Use the strconv.Atoi function
* to convert the string parameter into an integer. You can see its documentation with go doc strconv.Atoi.
 */
package main

import (
	"net/http"
	"log"
	"fmt"
	"strconv"
	"os"
)

func lissajousServer() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		cycles, err := strconv.Atoi(r.Form.Get("cycles"))

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			// just print it
			fmt.Println(cycles)
		}

		lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
