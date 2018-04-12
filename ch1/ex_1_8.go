/*
* Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL
* if it is missing. You might want to use strings.HasPrefix.
 */

package main

import (
	"fmt"
	"strings"
)

func autoPrefix() {
	urls := []string{"https://jmy5945hh.github.io/", "www.ukulelecn.com/forum.php"}
	prefix := "http://"

	for _, url := range urls {

		if false == strings.HasPrefix(url, prefix) {
			url = prefix + url
			fmt.Println(url)
		}
		//...
	}
}