package service

import (
	"os"
	"net/http"
	"io"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type ImageCreator interface {
	CreateImage() (bool, error)
}

type Image struct {
	Id       string
	Path     string
	Source   string
	ImageType        string
	InsertedDatetime time.Time
}

func NewImage(path, fileName, source string) (*Image) {
	return &Image{
		Path:     path,
		Id: fileName,
		Source:   source,
		ImageType: "jpg",
	}
}

func (i Image) CreateImage() (bool, error) {
	fileNameWithPath := validateFileInfo(i.Path, i.Id, i.ImageType)

	if fileNameWithPath == "" {
		return false, errors.New("Invalid file or path")
	}
	img, err := createBlankFile(fileNameWithPath)
	if err != nil {
		return false, err
	}
	defer img.Close()
	return getContentAndCopy(i.Source, img)
}
func validateFileInfo(path, fileName,fileTYpe string) (string) {
	fileNameWithPath := fmt.Sprintf("%s/%s.%s", path, fileName, fileTYpe)
	if validatePath(path) == false {
		logrus.Errorf("Invalid file path %s", path)
		return ""
	}
	if validateFileType(fileTYpe) == false {
		logrus.Errorf("Invalid file Name %s", fileName)
		return ""
	}
	return fileNameWithPath
}
func createBlankFile(fileName string) (*os.File, error) {
	img, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return img, nil
}
func getContentAndCopy(source string, img *os.File) (bool, error) {
	fmt.Sprintf("%v", source)
	resp, err := http.Get(source)
	if resp == nil || err != nil {
		logrus.Errorf("Unable to get the file from the url provided - %s", err.Error())
		return false, err
	}
	defer resp.Body.Close()

	fileSize, err := io.Copy(img, resp.Body)
	if err != nil {
		logrus.Errorf("Invalid response from the url - %s", err.Error())
		return false, err
	}

	if fileSize < 0 {
		logrus.Errorf("Error while creating the file - %s", err.Error())
		return false, errors.New("Error while creating the file")
	}
	return true, nil
}

func validateFileType(fileType string) (bool) {
	//TODO need to add other file types from config
	if fileType != "jpg" {
		return false
	}

	return true
}
func validatePath(path string) (bool) {
	stat, err := os.Stat(path)
	if err != nil {
		logrus.Errorf("Invalid file path %s", err.Error())
		return false
	}
	if stat.IsDir() == false {
		logrus.Error("Invalid file path not a directory")
		return false
	}

	return true
}
