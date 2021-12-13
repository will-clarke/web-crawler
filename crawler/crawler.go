package crawler

import (
	"fmt"
	"net/url"
	"time"
)

type webCrawler struct {
	url            string
	linkCountLimit int
	timeLimit      time.Duration
	ignoreErrs     bool
	store          Store
}

// TODO: this Store interface is a bit idealistic.
// For a real store we'd reality want to do some error-handling.
type Store interface {
	Get(string) bool
	Put(string)
}

func (c *webCrawler) CrawlWeb() error {
	rawURL := c.url
	url, err := url.Parse(rawURL)
	if err != nil {
		// TODO: better logging or whatever
		return err
	}
	links, err := c.GetLinksFromURL(url)
	fmt.Printf("links: %+v\n", links)
	fmt.Printf("err: %+v\n", err)
	return nil
}

func (c *webCrawler) GetLinksFromURL(u *url.URL) ([]*url.URL, error) {
	return nil, nil
}
