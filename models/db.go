package models

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/prog-image/middleware"
	"fmt"
)
const tableName = "images"
func InitDB(db middleware.Db) (*sql.DB, error) {
	//for mysql
	dbConnectionString := fmt.Sprintf("%s:%s@/%s", db.User, db.Password, db.Source)
	if db.Type == "sqlite3" {
		dbConnectionString = db.Source
	}
	dbConnector, err := sql.Open(db.Type, dbConnectionString)
	if err != nil {
		logrus.Errorf("Unable to start database %s", err.Error())
		return nil, err
	}
	return dbConnector,nil
}
func CreateTable(db *sql.DB) (error) {
	sql_table := `
	CREATE TABLE IF NOT EXISTS `+tableName+`(
		Id TEXT NOT NULL PRIMARY KEY,
		source TEXT,
		imageType TEXT,
		InsertedDatetime DATETIME
	);
	`
	_, err := db.Exec(sql_table)
	if err != nil {
		return err
	}
	return nil
}
