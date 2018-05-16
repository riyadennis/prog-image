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
	fileNameWithPath := ValidateFileInfo(i.Path, i.FileName)
	if fileNameWithPath == "" {
		return false, errors.New("Invalid file or path")
	}
	img, err := CreateBlankFile(fileNameWithPath)
	if err != nil {
		return false, err
	}
	defer img.Close()

	resp, err := http.Get(i.Source)
	defer resp.Body.Close()

	if err != nil {
		return false, err
	}

	fileSize, error := io.Copy(img, resp.Body)
	if fileSize > 0 {
		return true, nil
	}
	return false, error
}
func ValidateFileInfo(path, fileName string) (string) {
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
func CreateBlankFile(fileName string) (*os.File, error) {
	img, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return img, nil
}
