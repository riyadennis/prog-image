package cmd

import (
	"github.com/prog-image/middleware"
	"github.com/sirupsen/logrus"
	"github.com/prog-image/models"
)

func ExecuteCommand(migrateChoice string, config *middleware.Config) (bool) {
	db, err := models.InitDB(config.Prog.Db)
	if err != nil {
		logrus.Errorf("Unable initial  %s", err.Error())
	}
	if migrateChoice == "up"{
		return MigrateUp(db, config.Prog.Db.Source)

	}
	if migrateChoice == "down"{
		return MigrateDown(db, config.Prog.Db.Source)
	}
	return false
}
