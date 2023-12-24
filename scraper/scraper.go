// scraper/scraper.go
package scraper

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Quote represents a quote with text and author
type Quote struct {
	Text   string
	Author string
}

// ScrapeQuotes fetches quotes from the provided URL
func ScrapeQuotes(url string) ([]Quote, error) {
	quotes := []Quote{}

	// Fetch the HTML content of the page
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Find and extract quotes
	doc.Find(".quoteDetails").Each(func(i int, s *goquery.Selection) {
		text := cleanText(s.Find(".quoteText").Text())
		author := cleanAuthor(s.Find(".authorOrTitle").Text())

		// Create a Quote struct
		quote := Quote{
			Text:   text,
			Author: author,
		}

		// Append the quote to the quotes slice
		quotes = append(quotes, quote)
	})

	return quotes, nil
}

// Clean up quote text
func cleanText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "“", "")
	text = strings.ReplaceAll(text, "”", "")
	text = strings.ReplaceAll(text, "\n", "")

	parts := strings.Split(text, "―")
	if len(parts) > 0 {
		text = strings.TrimSpace(parts[0])
	}
	return text
}

// Clean up author text
func cleanAuthor(author string) string {
	author = strings.TrimSpace(author)
	author = strings.ReplaceAll(author, "\n", "")
	return author
}

// Tag represents a quote tag with a name and a link
type Tag struct {
	Name string
	Link string
}

// ScrapeTags fetches tags from the provided URL
func ScrapeTags(url string) ([]Tag, error) {
	tags := []Tag{}

	// Fetch the HTML content of the page
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Find and extract tags
	doc.Find(".listTagsTwoColumn li").Each(func(i int, s *goquery.Selection) {
		tagLink, exists := s.Find("a").Attr("href")
		if exists {
			tagName := strings.TrimSpace(s.Text())
			tag := Tag{
				Name: tagName,
				Link: tagLink,
			}

			// Append the tag to the tags slice
			tags = append(tags, tag)
		}
	})

	return cleanTag(tags, url), nil
}

func cleanTag(tags []Tag, url string) []Tag {
	for i, tag := range tags {
		tag.Link = strings.TrimPrefix(tag.Link, "/quotes/")
		tag.Link = url + tag.Link
		tags[i].Link = tag.Link

		parts := strings.Split(tag.Name, "\n")
		if len(parts) > 0 {
			tag.Name = strings.TrimSpace(parts[0])
		}
		tags[i].Name = tag.Name
	}
	return tags
}
