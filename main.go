package main

import (
	"fmt"
	"log"
	"quote-scraper/config"
	"quote-scraper/db"
	"quote-scraper/scraper"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}
	conn, err := db.OpenDB(*conf)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	url := "https://www.goodreads.com/quotes/"
	tags, err := scraper.ScrapeTags(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tags)

	for _, tag := range tags {
		for page_num := 1; page_num < 101; page_num++ {
			url := tag.Link + "?page=" + fmt.Sprint(page_num)
			quotes, err := scraper.ScrapeQuotes(url)
			if err != nil {
				log.Fatal(err)
			}
			db.InsertQuotes(conn, quotes, tag.Name)
		}
	}

}
