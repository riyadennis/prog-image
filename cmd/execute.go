package cmd

import (
	"github.com/prog-image/middleware"
	"github.com/sirupsen/logrus"
	"github.com/prog-image/models"
)

func ExecuteCommand(argsWithMain []string, config *middleware.Config) {
	db, err := models.InitDB(config.Prog.Db)
	if err != nil {
		logrus.Errorf("Unable initial  %s", err.Error())
	}
	if len(argsWithMain) > 1 {
		commandName := argsWithMain[1]
		subCommand := argsWithMain[2]
		if commandName == "migrate" && subCommand == "down" {
			MigrateDown(db, config.Prog.Db.Source)
			return
		}
		MigrateUp(db, config.Prog.Db.Source)
	}
	return
}
