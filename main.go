package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	url := "http://go-colly.org/"

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Title:", e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)
}
