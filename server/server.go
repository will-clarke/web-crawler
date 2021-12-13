package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	return r
}
