package main

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"net/http"
	"log"
	"github.com/prog-image/handlers"
	"github.com/prog-image/middleware"
)

func main() {
	Run(8080)
}

func Run(port int) {
	route := httprouter.New()
	route.GET("/healthcheck", handlers.HealthCheck)
	route.POST("/upload", handlers.UploadHandler)
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listenning to port %s \n", addr)
	log.Fatal(http.ListenAndServe(addr, middleware.ConfigMiddleWare(route)))
}
