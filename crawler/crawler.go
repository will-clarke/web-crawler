package crawler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type WebCrawler struct {
	InitialURL     string
	linkCountLimit int
	timeLimit      time.Duration
	ignoreErrs     bool
	urlStore       URLStore
	HttpClient     *http.Client
}

// TODO: this Store interface is a bit idealistic.
// For a real store we'd reality want to do some error-handling.
type URLStore interface {
	Get(string) bool
	Put(string)
	GetAllKeys() []string
}

func (c *WebCrawler) StartWebCrawl() ([]string, error) {
	err := c.Crawl(c.InitialURL)
	if err != nil {
		// TODO: better logging
		return nil, err
	}

	return c.urlStore.GetAllKeys(), nil
}

func (c *WebCrawler) alreadyCrawledPage(u string) bool {
	return c.urlStore.Get(u)
}

func (c *WebCrawler) Crawl(u string) error {
	time.Sleep(time.Millisecond * 500) // TODO: remove this... just tmp for testing so we don't go mental

	if c.alreadyCrawledPage(u) {
		return nil
	}

	links, err := c.GetLinksFromURL(u)
	fmt.Printf("links: %+v\n", links)
	if err != nil {
		// TODO: better logging
		return err
	}
	return nil
}

func (c *WebCrawler) GetLinksFromURL(u string) ([]string, error) {
	links := []string{}

	resp, err := c.HttpClient.Get(u)
	if err != nil {
		// TODO: maybe wrap error
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		a, _ := s.Attr("href")

		if a != "" {
			links = append(links, a)
		}

	})

	return links, nil
}

// IsInternalLink is a slightly hacky way of determining whether
// a given link is internal
func (c *WebCrawler) IsInternalLink(l string) bool {
	if strings.HasPrefix(l, "/") {
		return true // is likely a relative url
	}
	if strings.HasPrefix(l, c.InitialURL) {
		return true // starts with the same domain name
	}
	return false
}
