package main

import (
	"github.com/prog-image/middleware"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"os"
	"github.com/prog-image/cmd"
	"github.com/prog-image/handlers"
)

func main() {
	config, err := middleware.GetConfig()
	if err != nil {
		logrus.Errorf("Unable to fetch config %s", err.Error())
	}
	if len(os.Args) > 2 {
		cmd.ExecuteCommand(os.Args, config)
		os.Exit(0)
	}
	handlers.Run(config)
}
