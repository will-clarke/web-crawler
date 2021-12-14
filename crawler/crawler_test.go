package crawler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"git.sr.ht/~will-clarke/web-crawler/crawler"
	"github.com/stretchr/testify/assert"
)

func Test_webCrawler_GetLinksFromURL(t *testing.T) {
	tests := []struct {
		name          string
		htmlFilePath  string
		expectedLinks []string
	}{
		{
			name:          "it returns an error if it can't parse the file",
			htmlFilePath:  filepath.Join("fixtures", "nothing.txt"),
			expectedLinks: []string{},
		},
		{
			name:          "it returns no links for an empty file",
			htmlFilePath:  filepath.Join("fixtures", "nothing.txt"),
			expectedLinks: []string{},
		},
		{
			name:          "it returns no links for a web page with no links",
			htmlFilePath:  filepath.Join("fixtures", "no-links.html"),
			expectedLinks: []string{},
		},
		{
			name:          "it returns a sinle link if there's only a single link present",
			htmlFilePath:  filepath.Join("fixtures", "one-link.html"),
			expectedLinks: []string{"ONE LINK!!!"},
		},
		{
			name:         "it finds all the links from a web page",
			htmlFilePath: filepath.Join("fixtures", "example.html"),
			expectedLinks: []string{
				"/index.html",
				"/",
				"/posts.html",
				"/about.html",
				"/tags.html",
				"posts/2021-08-19--function-composition-is-super-cool.html",
				"/tags/haskell.html",
				"/tags/fp.html",
				"posts/2020-07-16--force-a-script-to-run-sudo.html",
				"https://git.sr.ht/~will-clarke/super-simple-static-site-generator",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := crawler.WebCrawler{
				HttpClient: http.DefaultClient,
			}

			stubbedServer := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					htmlBytes, err := ioutil.ReadFile(tt.htmlFilePath)
					assert.NoError(t, err)
					w.Write(htmlBytes)
				}),
			)
			defer stubbedServer.Close()

			links, err := c.GetLinksFromURL(stubbedServer.URL)
			assert.Equal(t, links, tt.expectedLinks)
			assert.NoError(t, err)
		})
	}
}

func TestWebCrawler_IsInternalLink(t *testing.T) {
	tests := []struct {
		name               string
		initialURL         string
		link               string
		expectedIsInternal bool
	}{
		{
			name:               "can identify internal paths from the initial domain",
			initialURL:         "https://example.com",
			link:               "https://example.com/obviously-internal",
			expectedIsInternal: true,
		},
		{
			name:               "can identify relative paths (which kind of by definition are internal)",
			initialURL:         "https://example.com",
			link:               "/obviously-internal",
			expectedIsInternal: true,
		},
		{
			name:               "will spot another domain",
			initialURL:         "https://example.com",
			link:               "https://example-NOT-THE-SAME-ONE.com",
			expectedIsInternal: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crawler.WebCrawler{
				InitialURL: tt.initialURL,
			}
			isInternal := c.IsInternalLink(tt.link)
			assert.Equal(t, tt.expectedIsInternal, isInternal)
		})
	}
}
