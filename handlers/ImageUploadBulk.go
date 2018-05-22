package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"github.com/prog-image/middleware"
	"github.com/sirupsen/logrus"
	"github.com/prog-image/service"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/pkg/errors"
)

type Uploader interface {
	Upload(filename, url string, config *middleware.Config, imageType string) (bool, error)
	GetFileName() string
}
type UploadHelper struct {
	FileName string
}
type UploadedImages struct {
	Images []*RequestImage
}

func UploadBulkHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.Body == nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uploadedImages := UploadedImages{}
	err = json.Unmarshal(requestBody, &uploadedImages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config, err := middleware.GetConfigFromContext(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := UploadHelper{}
	builkUploaded, err := BulkUpload(u, uploadedImages, config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if builkUploaded{
		res := createResponse("All images requestBody successfully", "Success", http.StatusOK)
		jsonResponseDecorator(res, w)
	}
}
func BulkUpload(u Uploader, images UploadedImages, config *middleware.Config) (bool, error){
	if len(images.Images) < 1{
		return false, errors.New("Invalid images")
	}
	for _, image := range images.Images{
		fileName := u.GetFileName()
		uploaded, err := u.Upload(fileName, image.Uri, config, image.ImageType)
		if err != nil {
			return false, err
		}
		err = SaveDataForTheImage(fileName, image, config)
		if err != nil {
			return false, nil
		}
		if uploaded {
			logrus.Infof("Successfully uploaded from url %s, with filename %s",image.Uri, fileName)
		}
		//clearing image struct
		image = nil
	}
	return true, nil

}
func(u UploadHelper) GetFileName() string{
	return fmt.Sprintf("%s", uuid.Must(uuid.NewV1(), nil))
}
func (u UploadHelper) Upload(filename,url string, config *middleware.Config, imageType string) (bool, error) {
	image := service.NewImage(filename, url, imageType)
	createdImage, err := image.CreateImage(config)
	if err != nil || createdImage == false{
		return false,err
	}
	return true, nil
}
