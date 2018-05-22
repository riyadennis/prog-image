package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"io/ioutil"
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
	u := UploadHelper{}
	filename := u.GetFileName()
	uploaded, err := u.Upload(filename, requestImage.Uri, config, requestImage.ImageType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if uploaded {
		err = SaveDataForTheImage(filename, &requestImage, config)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	detail := fmt.Sprintf("Image URL: %s", GetLocalImageURL(config, filename, requestImage.ImageType))
	res := createResponse(detail, "Success", http.StatusOK)
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


