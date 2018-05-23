package handlers

import (
	"testing"
	"github.com/prog-image/middleware"
	"github.com/stretchr/testify/assert"
	"os"
	"fmt"
	"net/http"
	"net/http/httptest"
	"github.com/julienschmidt/httprouter"

	_ "github.com/mattn/go-sqlite3"
)

const testImage = "testfile"

func TestCreateNewImageForInvalidImageType(t *testing.T) {
	config := middleware.Config{}
	err := CreateNewImageForImageType(testImage, "invaliderr", "png", &config)
	assert.Error(t, err)
}
func TestGetImageHandlerWithInvalidConfig(t *testing.T) {
	req, err := http.NewRequest("GET", "/images/sdd", nil)
	if err != nil {
		t.Fatal(err)
	}
	writer := httptest.NewRecorder()
	route := httprouter.New()
	route.GET("/images/:image_id", GetImageHandler)
	route.ServeHTTP(writer, req)
	assert.Equal(t, http.StatusInternalServerError, writer.Code)
}
func TestGetImageHandlerWithValidConfig(t *testing.T) {
	req, err := http.NewRequest("GET", "/images/sdd", nil)
	if err != nil {
		t.Fatal(err)
	}
	writer := httptest.NewRecorder()
	route := httprouter.New()
	context := ManageConfig(req)
	route.GET("/images/:image_id", GetImageHandler)
	route.ServeHTTP(writer, req.WithContext(context))
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}
func TestGetImageHandlerWithValidDatabaseConfig(t *testing.T) {
	req, err := http.NewRequest("GET", "/images/sdd", nil)
	if err != nil {
		t.Fatal(err)
	}
	writer := httptest.NewRecorder()
	route := httprouter.New()
	context := ManageConfig(req)
	route.GET("/images/:image_id", GetImageHandler)
	route.ServeHTTP(writer, req.WithContext(context))
	assert.Equal(t, http.StatusBadRequest, writer.Code)
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