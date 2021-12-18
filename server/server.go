package server

import (
	"net/http"

	"git.sr.ht/~will-clarke/web-crawler/crawler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TopSecretUsername = "user"
const TopSecretPassword = "pass"

type Servable interface {
	Run(...string) error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// URLStore is implemented by Store (../store/store.go)
type URLStore interface {
	Exists(string, string) bool
	Put(string, string)
	GetAllKeys(string) []string
}

type crawlRequest struct {
	Url string `json:"url" binding:"required"`
}

func SetupRouter(store URLStore) Servable {
	r := gin.Default()

	// some super basic basic auth
	auth := r.Group("/", gin.BasicAuth(gin.Accounts{
		TopSecretUsername: TopSecretPassword,
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
			"links": store.GetAllKeys(crawlID),
		})
	})

	auth.POST("/crawl", func(c *gin.Context) {
		id := uuid.New()
		var request crawlRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"id":    id.String(),
				"error": err})
			return
		}

		webCrawler, err := crawler.NewWebCrawler(store, id, request.Url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"id":    id.String(),
				"error": err})
			return
		}

		go webCrawler.Crawl(request.Url)

		c.JSON(http.StatusOK, gin.H{
			"id":    id.String(),
			"error": err,
		})
	})

	return r
}
