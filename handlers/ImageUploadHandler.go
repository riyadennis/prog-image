package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"io/ioutil"
	"github.com/prog-image/service"
	"github.com/satori/go.uuid"
	"fmt"
	"github.com/prog-image/middleware"
	"github.com/prog-image/models"
)

type RequestImage struct {
	Uri string `json:"uri"`
	ImageType string `json:"type"`
}
func UploadHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Body == nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestImage := RequestImage{}
	err = json.Unmarshal(requestBody, &requestImage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config, err := middleware.GetConfigFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	filename := fmt.Sprintf("%s", uuid.Must(uuid.NewV1(), nil))
	image := service.NewImage(filename, requestImage.Uri, requestImage.ImageType)
	created, err := image.CreateImage(config)
	if err != nil {
		fmt.Printf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if created == false {
		w.Write([]byte("Unable to save the image"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = SaveDataForTheImage(filename, &requestImage, config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	imageStruct, err := models.GetImage(filename, config.Prog.Db)
	responseDetail := fmt.Sprintf("Image Name: %s.%s",imageStruct.Id, image.ImageType)
	res := createResponse(responseDetail, "Success", http.StatusOK)
	jsonResponseDecorator(res, w)

}
func SaveDataForTheImage(fileName string,  requestImage *RequestImage , config *middleware.Config) (error){
	err :=  models.SaveImage(fileName, requestImage.Uri, config.Prog.Db)
	if err != nil {
		return err
	}
	err = models.SaveImageType(fileName, requestImage.ImageType, config.Prog.Db)
	if err != nil {
		return err
	}
	return nil
}


