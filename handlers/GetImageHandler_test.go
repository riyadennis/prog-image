package handlers

import (
	"testing"
	"github.com/prog-image/middleware"
	"github.com/stretchr/testify/assert"
	"fmt"
)
const testImage = "testfile"
func TestCreateNewImageForInvalidImageType(t *testing.T) {
	config := middleware.Config{}
	imageName, err := CreateNewImageForImageType(testImage, "invaliderr", "png", &config)
	assert.Error(t, err)
	assert.Empty(t, imageName)
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
	imageName, err := CreateNewImageForImageType(testImage, newFormat, "jpg", &config)
	expectedImageName := fmt.Sprintf("%s/%s.%s", config.Prog.Folder,testImage,  newFormat)
	assert.NoError(t, err)
	assert.Equal(t, imageName, expectedImageName)
}
