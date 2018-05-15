package service

import (
	"os"
	"net/http"
	"fmt"
	"io"
)

type ImageCreator interface {
	CreateImage() (bool, error)
}

type Image struct {
	Type   string
	Source string
	Path   string
}

func (i Image) CreateImage() (bool, error) {
	imageName := fmt.Sprintf("%s/imageName.%s", i.Path, i.Type)
	img, err := os.Create(imageName)
	if err != nil {
		return false, nil
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
