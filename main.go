package main

import "git.sr.ht/~will-clarke/web-crawler/server"

func main() {
	router := server.SetupRouter()
	router.Run()
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
