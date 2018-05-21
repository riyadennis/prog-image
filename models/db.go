package models

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/prog-image/middleware"
	"fmt"
	"time"
)

const tableName = "images"

type Image struct {
	Id               string
	Source           string
	ImageType        string
	InsertedDatetime time.Time
}

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
func SaveImage(filename, uri string, confDb middleware.Db) (error) {
	db, err := InitDB(confDb)
	if err != nil {
		logrus.Errorf("Unable to save image details %s", err.Error())
		return err
	}
	query := "INSERT INTO " + tableName + "(id,source,imageType) VALUES('" + filename + "', '" + uri + "', 'jpg')"
	res, err := db.Exec(query)
	if err != nil {
		logrus.Errorf("Unable to save image details %s", err.Error())
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected != 1 {
		logrus.Errorf("Unable to save image details %s", err.Error())
		return err
	}
	return nil
}

func GetImage(fileName string, confDb middleware.Db) (*Image, error) {
	var source, imageType string
	var image Image
	db, err := InitDB(confDb)
	if err != nil {
		logrus.Errorf("Unable to get image details %s", err.Error())
		return nil, err
	}
	query := "SELECT source,imageType from " + tableName + " where id = '" + fileName + "'"
	row := db.QueryRow(query)
	err = row.Scan(&source, &imageType)
	if err != nil {
		logrus.Errorf("Unable to get image details %s", err.Error())
		return nil, err
	}
	image.Source = source
	image.ImageType = imageType
	image.Id = fileName
	return &image, nil
}
