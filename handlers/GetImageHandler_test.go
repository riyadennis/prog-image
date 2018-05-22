package handlers

import (
	"testing"
	"github.com/prog-image/middleware"
	"github.com/stretchr/testify/assert"
	"os"
	"fmt"
)
const testImage = "testfile"
func TestCreateNewImageForInvalidImageType(t *testing.T) {
	config := middleware.Config{}
	err := CreateNewImageForImageType(testImage, "invaliderr", "png", &config)
	assert.Error(t, err)
}

func TestCreateNewImageForValidImageTypeWithInvalidOldFile(t *testing.T) {
	allowedFileTypes := make([]string, 2)
	allowedFileTypes[0] = "png"
	allowedFileTypes[1] = "jpg"
	prog := middleware.Prog{
		Folder: "../images",
		FileType:allowedFileTypes,
	}
	config := middleware.Config{
		Prog:prog,
	}
	newFormat := "png"
	newFilePath := fmt.Sprintf("%s/%s.%s", prog.Folder, testImage, newFormat)
	os.Remove(newFilePath)
	err := CreateNewImageForImageType(testImage, newFormat, "jpg", &config)
	assert.NoError(t, err)
}
