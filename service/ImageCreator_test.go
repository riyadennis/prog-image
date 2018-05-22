package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/prog-image/middleware"
)

func TestCreateImageWithInvalidDirectory(t *testing.T) {
	i := NewImage("testfile.png", "https://fyf.tac-cdn.net/images/products/large/BF116-11KM.jpg", "jpg")
	allowedFileTypes := make([]string, 1)
	allowedFileTypes[0] = "jpg"
	config := setUpConf(allowedFileTypes, "invalid")
	created, error := i.CreateImage(config)
	assert.Error(t, error)
	assert.Equal(t, false, created)
}

func TestCreateImageWithInvalidFileType(t *testing.T) {
	i := NewImage("testfile", "https://fyf.tac-cdn.net/images/products/large/BF116-11KM.jpg", "jpg")

	allowedFileTypes := make([]string, 1)
	allowedFileTypes[0] = "invalid"

	config := setUpConf(allowedFileTypes, "../images")
	created, error := i.CreateImage(config)
	assert.Error(t, error)
	assert.Equal(t, false, created)
}

func TestCreateImageWithValidFileName(t *testing.T) {
	i := NewImage("testfile.jpg", "https://fyf.tac-cdn.net/images/products/large/BF116-11KM.jpg", "jpg")
	allowedFileTypes := make([]string, 1)
	allowedFileTypes[0] = "jpg"
	config := setUpConf(allowedFileTypes, "../images")
	created, error := i.CreateImage(config)

	assert.NoError(t, error)
	assert.Equal(t, true, created)
}
func TestCreateImageWithInvalidSource(t *testing.T) {
	i := NewImage("testfile.jpg", "http:116-11KM.jpg", "jpg")
	allowedFileTypes := make([]string, 1)
	allowedFileTypes[0] = "jpg"

	config := setUpConf(allowedFileTypes, "../images")
	created, error := i.CreateImage(config)

	assert.Error(t, error)
	assert.Equal(t, false, created)
}
func setUpConf(allowedFileTypes []string, folder string) (*middleware.Config) {
	prog := middleware.Prog{
		Folder:   folder,
		FileType: allowedFileTypes,
	}
	return &middleware.Config{
		Prog: prog,
	}
}
