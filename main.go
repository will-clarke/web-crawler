package main

import (
	"log"

	"git.sr.ht/~will-clarke/web-crawler/server"
	"git.sr.ht/~will-clarke/web-crawler/store"
)

func main() {
	s := store.NewStore()
	router := server.SetupRouter(s)

	err := router.Run()
	if err != nil {
		log.Fatal("router errored", err)
	}
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
