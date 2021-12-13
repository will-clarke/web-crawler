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
	urlStore       URLStore
}

// TODO: this Store interface is a bit idealistic.
// For a real store we'd reality want to do some error-handling.
type URLStore interface {
	Get(string) bool
	Put(string)
	GetAllKeys() []string
}

func (c *webCrawler) CrawlWebStart() ([]string, error) {
	rawURL := c.url
	url, err := url.Parse(rawURL)
	if err != nil {
		// TODO: better logging or whatever
		return nil, err
	}
	err = c.CrawlWeb(url)
	if err != nil {
		// TODO: better logging
		return nil, err
	}

	return c.urlStore.GetAllKeys(), nil
}

func (c *webCrawler) alreadyCrawledPage(u *url.URL) bool {
	return c.urlStore.Get(u.String())
}

func (c *webCrawler) CrawlWeb(u *url.URL) error {
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

func (c *webCrawler) GetLinksFromURL(u *url.URL) ([]*url.URL, error) {
	return nil, nil
}
