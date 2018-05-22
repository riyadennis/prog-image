package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"github.com/prog-image/middleware"
	"github.com/sirupsen/logrus"
	"github.com/prog-image/service"
	"github.com/prog-image/models"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/pkg/errors"
)

type Uploader interface {
	Upload(filename, url, path string) (bool, error)
	GetFileName() string
}
type Uploaded struct {
	FileName string
}
type UploadedImages struct {
	Images []*UploadedImage
}

func UploadBulkHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.Body == nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	uploaded, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uploadedImages := UploadedImages{}
	err = json.Unmarshal(uploaded, &uploadedImages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config, err := middleware.GetConfigFromContext(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := Uploaded{}
	builkUploaded, err := BulkUpload(u, uploadedImages, config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if builkUploaded{
		res := createResponse("All images uploaded successfully", "Success", http.StatusOK)
		jsonResponseDecorator(res, w)
	}
}
func BulkUpload(u Uploader, images UploadedImages, config *middleware.Config) (bool, error){
	if len(images.Images) < 1{
		return false, errors.New("Invalid images")
	}
	for _, image := range images.Images{
		fileName := u.GetFileName()
		uploaded, err := u.Upload(fileName, image.Uri, config.Prog.Folder)
		if err != nil {
			return false, err
		}
		err =  models.SaveImage(fileName, image.Uri, config.Prog.Db)
		if err != nil {
			return false, err
		}
		if uploaded {
			logrus.Infof("Successfully uploaded from url %s, with filename %s",image.Uri, fileName)
		}
		//clearing image struct
		image = nil
	}
	return true, nil

}
func(u Uploaded) GetFileName() string{
	return fmt.Sprintf("%s", uuid.Must(uuid.NewV1(), nil))
}
func (u Uploaded) Upload(filename, url, path string) (bool, error) {
	image := service.NewImage(path, filename, url)
	createdImage, err := image.CreateImage()
	if err != nil || createdImage == false{
		return false,err
	}
	return true, nil
}
