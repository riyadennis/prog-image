package handlers

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"net/http"
	"github.com/sirupsen/logrus"
)

func Run(port int) {
	route := httprouter.New()
	route.GET("/healthcheck", HealthCheck)
	route.POST("/upload", UploadHandler)
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listenning to port %s \n", addr)
	logrus.Fatal(http.ListenAndServe(addr, route))
}
