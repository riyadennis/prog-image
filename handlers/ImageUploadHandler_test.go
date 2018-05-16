package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"strings"
	"github.com/julienschmidt/httprouter"
)

func TestUploadHandlerNoRequestBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/upload", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	route := httprouter.New()
	route.POST("/upload", UploadHandler)
	route.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUploadHandlerInvalidRequestBody(t *testing.T) {
	body := strings.NewReader("Invalid body")
	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	route := httprouter.New()
	route.POST("/upload", UploadHandler)
	route.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
func TestUploadHandlerInValidJsonRequestBody(t *testing.T) {
	json := `{
	"uri": "http-beautiful-beauty.jpg"
}`
	body := strings.NewReader(json)
	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	route := httprouter.New()
	route.POST("/upload", UploadHandler)
	route.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
func TestUploadHandlerValidRequestBody(t *testing.T) {
	json := `{
	"uri": "https://images.pexels.com/photos/36764/marguerite-daisy-beautiful-beauty.jpg"
}`
	body := strings.NewReader(json)
	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	route := httprouter.New()
	route.POST("/upload", UploadHandler)
	route.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}