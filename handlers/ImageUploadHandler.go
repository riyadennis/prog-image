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

type UploadedImage struct {
	Uri string `json:"uri"`
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
	uploaded, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uploadedImage := UploadedImage{}
	err = json.Unmarshal(uploaded, &uploadedImage)
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
	image := service.NewImage(config.Prog.Folder, filename, uploadedImage.Uri)
	created, err := image.CreateImage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if created == false {
		w.Write([]byte("Unable to save the image"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = SaveImage(filename, uploadedImage.Uri, config.Prog.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	imageStruct, err := models.GetImage(filename, config.Prog.Db)
	responseDetail := fmt.Sprintf("Image Name: %s.%s",imageStruct.Id, image.ImageType)
	res := createResponse(responseDetail, "Success", http.StatusOK)
	jsonResponseDecorator(res, w)

}
func SaveImage(filename, uri string, configDb middleware.Db) (error){
	err :=  models.SaveImage(filename, uri, configDb)
	if err != nil {
		return err
	}
	return nil
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
