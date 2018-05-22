package service

import (
	"os"
	"net/http"
	"io"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
	"github.com/prog-image/middleware"
)

type ImageCreator interface {
	CreateImage() (bool, error)
}

type Image struct {
	Id       string
	Source   string
	ImageType        string
	InsertedDatetime time.Time
}

func NewImage(fileName, source, imageType string) (*Image) {
	return &Image{
		Id: fileName,
		Source:   source,
		ImageType: imageType,
	}
}

func (i Image) CreateImage(config *middleware.Config) (bool, error) {
	fileNameWithPath := validateFileInfo(config.Prog.Folder, i.Id, i.ImageType, config)

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
func validateFileInfo(path, fileName,fileType string, config *middleware.Config) (string) {
	fileNameWithPath := fmt.Sprintf("%s/%s.%s", path, fileName, fileType)
	if validatePath(path) == false {
		logrus.Errorf("Invalid file path %s", path)
		return ""
	}
	if ValidateFileType(fileType, config) == false {
		logrus.Errorf("Invalid file type %s", fileName)
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

func ValidateFileType(fileType string, config *middleware.Config) (bool) {
	allowedTypes := config.Prog.FileType
	for _, t := range allowedTypes{
		if fileType == t {
			return true
		}
	}
	return false
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
