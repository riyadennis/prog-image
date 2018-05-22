package models

import (
	"github.com/sirupsen/logrus"
	"github.com/prog-image/middleware"
	"strings"
)

const imageTypeTableName = "image_types"

type ImageType struct {
	id        string
	imageType string
}

func SaveImageType(imageId, imageType string, confDb middleware.Db) (error) {
	db, err := InitDB(confDb)
	if err != nil {
		logrus.Errorf("Unable to save image details %s", err.Error())
		return err
	}

	query := "INSERT INTO " + imageTypeTableName + "(id,image_type,InsertedDatetime) VALUES('" + strings.TrimSpace(imageId) + "', '" + strings.TrimSpace(imageType) + "' , '" + getCurrentTimeStamp() + "')"
	res, err := db.Exec(query)
	if err != nil {
		logrus.Errorf("Unable to save image type details %s", err.Error())
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected != 1 {
		logrus.Errorf("Unable to save image type details %s", err.Error())
		return err
	}
	return nil
}
func GetImageType(imageId string, confDb middleware.Db) (string, error) {
	var imageType string
	db, err := InitDB(confDb)
	if err != nil {
		logrus.Errorf("Unable to get image details %s", err.Error())
		return "", err
	}

	query := "SELECT image_type from " + imageTypeTableName + " where id = '" + strings.TrimSpace(imageId) + "'"
	row := db.QueryRow(query)
	err = row.Scan(&imageType)
	if err != nil {
		logrus.Errorf("Unable to get image type %s", err.Error())
		return "", err
	}
	return imageType, nil
}
