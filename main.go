package main

import (
	"test/router"
)

func main() {
	router := router.InitRouter()
	router.Run(":9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
