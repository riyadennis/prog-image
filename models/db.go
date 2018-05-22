package models

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/prog-image/middleware"
	"fmt"
	"time"
)

func InitDB(db middleware.Db) (*sql.DB, error) {
	//for mysql
	dbConnectionString := fmt.Sprintf("%s:%s@/%s?multiStatements=true", db.User, db.Password, db.Source)
	if db.Type == "sqlite3" {
		dbConnectionString = db.Source
	}
	dbConnector, err := sql.Open(db.Type, dbConnectionString)
	if err != nil {
		logrus.Errorf("Unable to start database %s", err.Error())
		return nil, err
	}
	return dbConnector, nil
}

func getCurrentTimeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}