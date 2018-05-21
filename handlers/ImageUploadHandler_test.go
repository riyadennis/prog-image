package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"strings"
	"github.com/julienschmidt/httprouter"
	 "context"
	"github.com/prog-image/middleware"
	"github.com/prog-image/models"

	_ "github.com/mattn/go-sqlite3"

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
	route.ServeHTTP(rr, req.WithContext(ManageConfig(req)))
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
	ctx := ManageConfig(req)
	route.ServeHTTP(rr, req.WithContext(ctx))
	assert.Equal(t, http.StatusOK, rr.Code)
}
func ManageConfig(req *http.Request) (context.Context){
	db := middleware.Db{Source: "test.db", Type: "sqlite3"}
	prog := middleware.Prog{
		Port: 8080,
		Folder: "../images",
		Db: db,
	}
	config := middleware.Config{
		Prog: prog,
	}
	dbConnec, _ := models.InitDB(db)
	statement, _ := dbConnec.Prepare("CREATE TABLE IF NOT EXISTS  images(id varchar(100) NOT NULL PRIMARY KEY,source varchar(100),imageType varchar(200),InsertedDatetime DATETIME);")
	statement.Exec()

	return context.WithValue(req.Context(), middleware.ContextKey, config)
}
func TestSaveImage(t *testing.T) {

}