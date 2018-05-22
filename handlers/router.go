package handlers

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"net/http"
	"github.com/prog-image/middleware"
	"github.com/sirupsen/logrus"
)

func Run(config *middleware.Config) {
	route := httprouter.New()
	route.GET("/healthcheck", HealthCheck)
	route.POST("/upload", UploadHandler)
	route.POST("/upload/bulk", UploadBulkHandler)

	fileConf := &FileConf{
		path: "images",
	}
	route.NotFound = fileConf.FileServer

	addr := fmt.Sprintf(":%d", config.Prog.Port)
	fmt.Printf("Listenning to port %s \n", addr)
	logrus.Fatal(http.ListenAndServe(addr, middleware.ConfigMiddleWare(route)))
}
