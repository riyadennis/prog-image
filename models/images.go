package models

import (
	"github.com/prog-image/middleware"
	"github.com/prog-image/service"
)

const tableName = "images"

func SaveImage(filename, uri string, confDb middleware.Db) (error) {
	db, err := InitDB(confDb)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO " + tableName + "(id,source,InsertedDatetime) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(filename, uri,getCurrentTimeStamp())
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected != 1 {
		return err
	}
	return nil
}

func GetImage(fileName string, confDb middleware.Db) (*service.Image, error) {
	var source string
	var image service.Image
	db, err := InitDB(confDb)
	if err != nil {
		return nil, err
	}
	query := "SELECT source from " + tableName + " where id = '" + fileName + "'"
	row := db.QueryRow(query)
	err = row.Scan(&source)
	if err != nil {
		return nil, err
	}
	image.Source = source
	image.Id = fileName
	return &image, nil
}
