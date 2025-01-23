package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

const UnknownStatusCode = 0

func CheckURL(url string) (string, error) {
	client := http.DefaultClient
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	return response.Request.Host, nil
}

func Scrape(url, host string) {
	client := http.DefaultClient

	// This check helps to avoid infinite loops (in the future, when I try to make this recursive)
	// we only check a link if it is not in visitedLinks
	// if !slices.Contains(visitedLinks, url) {

	body, statusCode, err := CheckStatus(url, client)
	if err != nil {
		log.Printf("%s\tError while checking link\t%d\t%v\n", url, statusCode, err)
		return
	}
	if statusCode != 200 {
		log.Printf("%s\tDead link\t%d\n", url, statusCode)
		return
	}
	fmt.Printf("%s\tValid link\t%d\n", url, statusCode)
	// TODO get links inside current site
	links, err := GetLinks(body)
	if err != nil {
		log.Printf("Error getting links: %v\n", err)
	}

	for _, link := range links {
		fmt.Println(link)
	}
}

// TODO I don' t like how this is going
func GetLinks(body []byte) ([]string, error) {
	links := []string{}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var extractLinks func(*html.Node)
	extractLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractLinks(c)
		}
	}

	extractLinks(doc)
	return links, nil
}

func CheckStatus(url string, client *http.Client) ([]byte, int, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, UnknownStatusCode, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}

	return body, response.StatusCode, nil
}
