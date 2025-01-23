package main

import (
	"flag"
	"fmt"
)

func main() {
	urlPtr := flag.String("url", "", "The URL to scrap")
	flag.Parse()

	if *urlPtr == "" {
		fmt.Println("Usage: ./scraper -url [site]")
		return
	}

	fmt.Println("URL is: " + *urlPtr)

	// TODO
	// 1. Create a client, can it be reusable?
	// 2. Connect to specified URL
	// 3. Store client and URL, this is what we need to start working
	//
	// Call scrape function and store valid links
	// Or dead links instead?
	// CheckStatus(*urlPtr)
	Scrape(*urlPtr)
}
