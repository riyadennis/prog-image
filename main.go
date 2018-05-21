package main

import (
	"github.com/prog-image/middleware"
	"context"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"os"
	"github.com/prog-image/cmd"
	"github.com/prog-image/handlers"
)

func main() {
	config, err := middleware.GetConfig(context.Background())
	if err != nil {
		logrus.Errorf("Unable to fetch config %s", err.Error())
	}
	cmd.ExecuteCommand(os.Args, config)
	handlers.Run(config.Prog.Port)
}
