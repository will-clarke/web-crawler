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

type WebCrawler struct {
	UrlStore   URLStore
	ID         uuid.UUID
	host       string
	scheme     string
	httpClient *http.Client
}

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

// TODO: this Store interface is a bit idealistic.
// For a real store we'd reality want to do some error-handling.
type URLStore interface {
	Get(string, string) bool
	Put(string, string)
	GetAllKeys(string) []string
}

func (c *WebCrawler) alreadyCrawledPage(s string) bool {
	return c.UrlStore.Get(c.ID.String(), s)
}

func (c *WebCrawler) Crawl(s string) {
	time.Sleep(time.Millisecond * 500) // TODO: remove this... just tmp for testing so we don't go mental

	if c.alreadyCrawledPage(s) {
		return
	}

	// we're telling our store that we've
	// visited this page
	c.UrlStore.Put(c.ID.String(), s)
	links, err := c.GetLinksFromURL(s)
	if err != nil {
		fmt.Printf("error getting links from url: %s\n", err.Error())
		return
	}
	for _, link := range links {
		if c.IsInternalLink(link) {
			go c.Crawl(link)
		}
	}
}

func (c *WebCrawler) GetLinksFromURL(s string) ([]string, error) {
	links := []string{}

	req, err := http.NewRequest(http.MethodGet, s, nil)
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

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

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
	if strings.Contains(l, c.host) {
		return true // starts with the same domain name
	}
	return false
}
