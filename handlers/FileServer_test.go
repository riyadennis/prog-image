package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)


func TestFileServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/testfile.jpg", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	route := httprouter.New()
	fileConf := &FileConf{
		path: "../images",
	}
	route.NotFound = fileConf.FileServer
	route.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}