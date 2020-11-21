package main

import "test/router"

func main() {
	router := router.InitRouter()
	router.Run(":9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//router := gin.Default()
	//
	//s := &http.Server{
	//	Addr:           ":9999",
	//	Handler:        router,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()
}
