package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Servable interface {
	Run(...string) error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

func SetupRouter() Servable {
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

	r.POST("/crawl", func(c *gin.Context) {
		id := uuid.NewString()
		c.JSON(200, gin.H{
			"randomID": id,
		})
	})

	return r
}
