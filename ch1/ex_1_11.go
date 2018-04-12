/*
* Exercise 1.11: Try fetchall with longer argument lists, such as samples from
* the top million websites available at alexa.com.How does the program behave if
* a website just doesn’t respond?(Section8.9 describes mechanisms for coping in such cases.)
 */

package main

import (
	"fmt"
	"time"
	"net/http"
	"os"
	"github.com/PuerkitoBio/goquery"
)

func parseTop50() []string {
	url := "https://www.alexa.com/topsites"
	resp, err := http.Get(url)

	var ret_urls []string

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return nil
	}

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		site := s.Find("a").Text()
		if len(site) > 0 {
			ret_urls = append(ret_urls, "http://"+site)
		}
	})

	return ret_urls
}

func fetchTop() {
	start := time.Now()
	ch := make(chan string)

	urls := parseTop50()

	fmt.Println(urls)

	for _, url := range urls {
		go fetchConCurr(url, ch) // start a goroutine
	}

	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}