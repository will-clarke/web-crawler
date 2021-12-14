package main

import (
	"log"

	"git.sr.ht/~will-clarke/web-crawler/server"
)

func main() {
	router := server.SetupRouter()
	err := router.Run()
	if err != nil {
		log.Fatal("router errored", err)
	}
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
