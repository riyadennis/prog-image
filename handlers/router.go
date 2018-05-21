package handlers

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/prog-image/middleware"
)

func Run(config *middleware.Config) {
	route := httprouter.New()
	route.GET("/healthcheck", HealthCheck)
	route.POST("/upload", UploadHandler)
	addr := fmt.Sprintf(":%d", config.Prog.Port)
	fmt.Printf("Listenning to port %s \n", addr)
	logrus.Fatal(http.ListenAndServe(addr, middleware.ConfigMiddleWare(route)))
}
