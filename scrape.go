package main

import (
	"fmt"
	"log"
	"net/http"
)

const UnknownStatusCode = 0

func Scrape(url string) {
	client := http.DefaultClient
	// pendingLinks := []string

	valid, statusCode, err := CheckStatus(url, client)
	if err != nil || valid != true {
		log.Printf("%s\tDead link\t%d\t%v\n", url, statusCode, err)
		return
	}
	// get links inside current site
	fmt.Printf("%s\tValid link\t%d\n", url, statusCode)
}

func CheckStatus(url string, client *http.Client) (bool, int, error) {
	response, err := client.Get(url)
	if err != nil {
		return false, UnknownStatusCode, err
	}
	defer response.Body.Close()

	return response.StatusCode == 200, response.StatusCode, nil
}
