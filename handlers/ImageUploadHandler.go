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
type ApiResponse struct {
	Status int
	Detail string
	Title  string
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
	//TODO check the image type from request with the allowed types
	filename := fmt.Sprintf("%s", uuid.Must(uuid.NewV1(), nil))
	image := service.NewImage(config.Prog.Folder, filename, requestImage.Uri, requestImage.ImageType)
	created, err := image.CreateImage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if created == false {
		w.Write([]byte("Unable to save the image"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err =  models.SaveImage(filename, requestImage.Uri, config.Prog.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	imageStruct, err := models.GetImage(filename, config.Prog.Db)
	responseDetail := fmt.Sprintf("Image Name: %s.%s",imageStruct.Id, image.ImageType)
	res := createResponse(responseDetail, "Success", http.StatusOK)
	jsonResponseDecorator(res, w)

}

func jsonResponseDecorator(response *ApiResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), response.Status)
		return
	}
}
func createResponse(detail, title string, status int) *ApiResponse {
	return &ApiResponse{
		Status: status,
		Detail: detail,
		Title:  title,
	}
}
