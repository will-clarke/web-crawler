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
	Get(string, string) bool
	Put(string, string)
	GetAllKeys(string) []string
}

type crawlRequest struct {
	Url string `json:"url" binding:"required"`
}

func SetupRouter(s URLStore) Servable {
	r := gin.Default()

	// some super basic basic auth
	auth := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user": "pass", // top secret credentials
	}))

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	auth.GET("/crawl/:id", func(c *gin.Context) {
		crawlID := c.Param("id")
		c.JSON(200, gin.H{
			"id":    crawlID,
			"links": s.GetAllKeys(crawlID),
		})
	})

	auth.POST("/crawl", func(c *gin.Context) {
		var request crawlRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := uuid.New()

		webCrawler, err := crawler.NewWebCrawler(s, id, request.Url)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}

		go webCrawler.Crawl(request.Url)

		c.JSON(200, gin.H{
			"id":    id.String(),
			"error": err,
			// probably should't be exposing our error messages...
		})
	})

	return r
}
