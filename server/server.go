package server

import (
	"net/http"

	"git.sr.ht/~will-clarke/web-crawler/crawler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Servable interface {
	Run(...string) error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type URLStore interface {
	Get(string) bool
	Put(string)
	GetAllKeys() []string
}

func SetupRouter(s URLStore) Servable {
	r := gin.Default()
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/crawl/:id", func(c *gin.Context) {
		crawlID := c.Param("id")
		c.JSON(200, gin.H{
			"crawlID": crawlID,
		})
	})

	r.POST("/crawl/:url", func(c *gin.Context) {
		url := c.Param("url")
		id := uuid.New()

		webCrawler := crawler.WebCrawler{
			InitialURL: url,
			UrlStore:   s,
			HttpClient: http.DefaultClient,
			ID:         id,
		}
		err := webCrawler.StartWebCrawl()

		c.JSON(200, gin.H{
			"id":  id.String(),
			"err": err,
			// probably should't be exposing our error messages...
		})
	})

	return r
}
