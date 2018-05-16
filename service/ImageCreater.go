package service

import (
	"os"
	"net/http"
	"io"
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"github.com/sirupsen/logrus"
)

type ImageCreator interface {
	CreateImage() (bool, error)
}

type Image struct {
	Path     string
	FileName string
	Source   string
}

func NewImage(path, fileName, source string) (*Image) {
	return &Image{
		Path:     path,
		FileName: fileName,
		Source:   source,
	}
}

func (i Image) CreateImage() (bool, error) {
	fileNameWithPath := validateFileInfo(i.Path, i.FileName)
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
func validateFileInfo(path, fileName string) (string) {
	fileNameWithPath := fmt.Sprintf("%s/%s", path, fileName)
	if validatePath(path) == false {
		logrus.Errorf("Invalid file path %s", path)
		return ""
	}
	if validateFileName(fileName) == false {
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

func validateFileName(fileName string) (bool) {
	fileExt := strings.LastIndex(fileName, ".")
	if fileExt < 0 {
		return false
	}
	fileExtension := fileName[fileExt+1:len(fileName)]
	if fileExtension != "jpg" {
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
