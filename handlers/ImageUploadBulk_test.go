package handlers

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/mock"
	"github.com/prog-image/middleware"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func TestUploadBulkHandlerWithNoRequestBody(t *testing.T) {
	request, err := http.NewRequest("POST", "/upload/bulk", nil)
	assert.NoError(t, err)
	writer := httptest.NewRecorder()
	UploadBulkHandler(writer, request, nil)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}
func TestUploadBulkHandlerWithInvalidRequestBody(t *testing.T) {
	requestBody := strings.NewReader("Invalid")
	request, err := http.NewRequest("POST", "/upload/bulk", requestBody)
	assert.NoError(t, err)
	writer := httptest.NewRecorder()
	UploadBulkHandler(writer, request, nil)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}
func TestUploadBulkHandlerWithValidRequestBody(t *testing.T) {
	//this is how your request body should look like
	jsonBody := `{
	"images":
	[
		{
			"uri": "https://images.pexels.com/photos/36764/marguerite-daisy-beautiful-beauty.jpg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260"
		},
		{
			"uri": "https://cdn.flowercompany.ca/wp-content/uploads/2017/01/My-Heart-to-Yours-497x600.jpg"
		}
	]
}
`
	requestBody := strings.NewReader(jsonBody)
	request, err := http.NewRequest("POST", "/upload/bulk", requestBody)
	assert.NoError(t, err)
	writer := httptest.NewRecorder()

	ctx := ManageConfig(request)

	router := httprouter.New()
	router.POST("/upload/bulk", UploadBulkHandler)
	router.ServeHTTP(writer, request.WithContext(ctx))
	// I can unmarshal my request and create a struct
	assert.Equal(t, writer.Code, http.StatusOK)
}
func TestUploaded_UploadInvalidFileAndPath(t *testing.T) {
	u := Uploaded{}
	uploaded, err := u.Upload("test", "test", "invalid")
	assert.Error(t, err)
	assert.False(t, uploaded)
}

type MockUploader struct {
	mock.Mock
	FileName string
}

func (m MockUploader) Upload(filename, url, path string) (bool, error) {
	args := m.Called(filename, url, path)
	return args.Bool(0), args.Error(1)
}
func (m MockUploader) GetFileName() (string) {
	args := m.Called()
	return args.String(0)
}
func TestBulkUploadWithInvalidImagesSlice(t *testing.T) {
	m := MockUploader{}
	images := UploadedImages{}
	conf := &middleware.Config{}
	uploaded, err := BulkUpload(m, images, conf)
	assert.False(t, uploaded)
	assert.Error(t, err)
}
func TestBulkUploadWithValidImageSlice(t *testing.T) {
	uImage := make([]*UploadedImage, 1)
	uImage[0] = &UploadedImage{
		Uri: "https://cdn.flowercompany.ca/wp-content/uploads/2017/01/My-Heart-to-Yours-497x600.jpg",
	}
	images := UploadedImages{Images: uImage}
	m := MockUploader{}
	m.FileName = "testfile"
	m.On("Upload", "testfile", uImage[0].Uri, "path").Return(true, nil)
	m.On("GetFileName").Return("testfile")
	db := middleware.Db{Source: testDbName, Type: "sqlite3"}
	conf := &middleware.Config{
		Prog: middleware.Prog{
			Folder: "path",
			Db:     db,
		},
	}
	setUpTestDB(db)
	uploaded, err := BulkUpload(m, images, conf)
	os.Remove(testDbName)
	assert.True(t, uploaded)
	assert.NoError(t, err)
}
