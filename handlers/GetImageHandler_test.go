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
		Folder:   "../images",
		FileType: allowedFileTypes,
	}
	config := middleware.Config{
		Prog: prog,
	}
	newFormat := "png"
	newFilePath := fmt.Sprintf("%s/%s.%s", prog.Folder, testImage, newFormat)
	os.Remove(newFilePath)
	err := CreateNewImageForImageType(testImage, newFormat, "jpg", &config)
	assert.NoError(t, err)
}
func TestFileConverter_ConvertFileFromJpgToPng(t *testing.T) {
	newFileType := "png"
	newFile, _ := os.Create("testfile.png")
	oldFileType := "jpg"
	oldFIlePath := fmt.Sprintf("../images/%s.%s", testImage, oldFileType)
	oldFile, _ := os.Open(oldFIlePath)
	fc := FileConverter{
		NewFileType: newFileType,
		NewFile:     newFile,
		OldFileType: oldFileType,
		OldFile:     oldFile,
	}
	err := fc.ConvertFileFromOnTypeToAnother()
	os.Remove("testfile.png")
	assert.NoError(t, err)
}
func TestFileConverter_ConvertFileFromPngToGif(t *testing.T) {
	newFileType := "gif"
	newFile, _ := os.Create("testfile.gif")
	oldFileType := "png"
	oldFIlePath := fmt.Sprintf("../images/%s.%s", testImage, oldFileType)
	oldFile, _ := os.Open(oldFIlePath)
	fc := FileConverter{
		NewFileType: newFileType,
		NewFile:     newFile,
		OldFileType: oldFileType,
		OldFile:     oldFile,
	}
	err := fc.ConvertFileFromOnTypeToAnother()
	os.Remove("testfile.gif")
	assert.NoError(t, err)
}
func TestFileConverter_ConvertFileFromPngToInvalid(t *testing.T) {
	newFileType := "jpegs"
	newFile, _ := os.Create("testfile.jpegs")
	oldFileType := "png"
	oldFIlePath := fmt.Sprintf("../images/%s.%s", testImage, oldFileType)
	oldFile, _ := os.Open(oldFIlePath)
	fc := FileConverter{
		NewFileType: newFileType,
		NewFile:     newFile,
		OldFileType: oldFileType,
		OldFile:     oldFile,
	}
	err := fc.ConvertFileFromOnTypeToAnother()
	os.Remove("testfile.jpegs")
	assert.Error(t, err)
}