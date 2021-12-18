package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
)

// WebCrawler is a struct that will crawl any *internal* links
// once we call `Crawl`.
type WebCrawler struct {
	UrlStore   URLStore
	ID         uuid.UUID
	host       string
	scheme     string
	httpClient *http.Client
}

// URLStore is an interface that `Store` implements (../store/store.go).
// We need this to keep track of which URLs we've already visited.
type URLStore interface {
	Exists(string, string) bool
	Put(string, string)
	GetAllKeys(string) []string
}

// NewWebCrawler is a constructor for WebCrawler.
// It has some logic to work out the host/scheme
// of a URL, which is useful for crawling relative links.
func NewWebCrawler(urlStore URLStore, id uuid.UUID, initialURL string) (*WebCrawler, error) {
	parsedURL, err := url.Parse(initialURL)
	if err != nil {
		return nil, err
	}
	return &WebCrawler{
		UrlStore:   urlStore,
		ID:         id,
		host:       parsedURL.Host,
		scheme:     parsedURL.Scheme,
		httpClient: http.DefaultClient,
	}, nil
}

// Crawl is the main function in this module. It recursively
// crawls through all internal links.
func (c *WebCrawler) Crawl(internalURL string) {
	time.Sleep(time.Millisecond * 50)

	// return early if we've crawled the page. We don't want to
	// waste time reading the same page and we could end up in an
	// infinite loop :scream:
	if c.alreadyCrawledPage(internalURL) {
		return
	}

	// we're telling our store that we've visited this page
	c.UrlStore.Put(c.ID.String(), internalURL)

	// we're parsing the HTML from the given url
	links, err := c.GetLinksFromURL(internalURL)
	if err != nil {
		fmt.Printf("error getting links from url: %s\n", err.Error())
		return
	}

	// for every internal link, call this `crawl` function again
	for _, link := range links {
		if c.IsInternalLink(link) {
			go c.Crawl(link)
		}
	}
}

// GetLinksFromURL fetches the HTML for a given URL, parses it and then
// extracts all the links.
func (c *WebCrawler) GetLinksFromURL(internalURL string) ([]string, error) {
	links := []string{}

	req, err := http.NewRequest(http.MethodGet, internalURL, nil)
	if err != nil {
		return nil, err
	}

	// some cases to cover any relative links
	req.URL.Scheme = c.scheme
	req.URL.Host = c.host
	req.Host = c.host

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bad status code: %d %s", resp.StatusCode, resp.Status)
	}

	// we're using a fab 3rd party HTML parser to get all the links
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("a").Each(func(i int, internalURL *goquery.Selection) {
		a, _ := internalURL.Attr("href")

		if a != "" {
			links = append(links, a)
		}

	})

	return links, nil
}

// IsInternalLink is a slightly hacky way of determining whether
// a given link is internal.
func (c *WebCrawler) IsInternalLink(s string) bool {
	if strings.HasPrefix(s, "/") {
		return true // is likely a relative url
	}
	if strings.Contains(s, c.host) {
		return true // starts with the same domain name
	}
	return false
}

// alreadyCrawledPage uses our store to work out if we've already
// crawled a given URL.
func (c *WebCrawler) alreadyCrawledPage(internalURL string) bool {
	return c.UrlStore.Exists(c.ID.String(), internalURL)
}
