package models

import (
	"github.com/prog-image/middleware"
	"github.com/sirupsen/logrus"
	"github.com/prog-image/service"
	"time"
)

const tableName = "images"

func SaveImage(filename, uri string, confDb middleware.Db) (error) {
	db, err := InitDB(confDb)
	if err != nil {
		logrus.Errorf("Unable to save image details %s", err.Error())
		return err
	}
	ts := time.Now().Format("2006-01-02 15:04:05")
	query := "INSERT INTO " + tableName + "(id,source,InsertedDatetime) VALUES('" + filename + "', '" + uri + "' , '"+ts+"')"
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

func GetImage(fileName string, confDb middleware.Db) (*service.Image, error) {
	var source string
	var image service.Image
	db, err := InitDB(confDb)
	if err != nil {
		logrus.Errorf("Unable to get image details %s", err.Error())
		return nil, err
	}
	query := "SELECT source from " + tableName + " where id = '" + fileName + "'"
	row := db.QueryRow(query)
	err = row.Scan(&source)
	if err != nil {
		logrus.Errorf("Unable to get image details %s", err.Error())
		return nil, err
	}
	image.Source = source
	image.Id = fileName
	return &image, nil
}
