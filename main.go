package main

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"net/http"
	"log"
	"github.com/prog-image/handlers"
	"github.com/prog-image/middleware"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/prog-image/models"
)

func main() {
	config, err := middleware.GetConfig(context.Background())
	if err != nil {
		logrus.Errorf("Unable to fetch config %s", err.Error())
	}
	db, err := models.InitDB(config.Prog.Db)
	if err != nil {
		logrus.Errorf("Unable initial  %s", err.Error())
	}
	err = models.CreateTable(db)
	if err != nil {
		logrus.Errorf("Unable set up database %s", err.Error())
	}
	Run(config.Prog.Port)
}

func Run(port int) {
	route := httprouter.New()
	route.GET("/healthcheck", handlers.HealthCheck)
	route.POST("/upload", handlers.UploadHandler)
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listenning to port %s \n", addr)
	log.Fatal(http.ListenAndServe(addr, middleware.ConfigMiddleWare(route)))
}
