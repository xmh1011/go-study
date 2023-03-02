package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

// Extract makes an HTTP GET request to the specified URL,
func Extract(url string) ([]string, error) {
	s := []string{}
	return s, nil
}

// Crawl the web breadth-first,
// starting from the command-line arguments.
func crawl(url string) []string {
	// ...
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}
