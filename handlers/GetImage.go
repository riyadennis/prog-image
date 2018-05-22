package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/prog-image/middleware"
	"github.com/prog-image/models"
	"github.com/prog-image/service"
)

func GetImageHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestImageId := params.ByName("image_id")
	requestImageType := req.URL.Query().Get("type")
	if requestImageType != ""{
		config, err := middleware.GetConfigFromContext(req.Context())
		if err != nil {
			http.Error(w, "Unable to process the request", http.StatusInternalServerError)
			return
		}
		if service.ValidateFileType(requestImageType, config) {

		}
	}
	config, err := middleware.GetConfigFromContext(req.Context())
	if err != nil {
		http.Error(w, "Unable to process the request", http.StatusInternalServerError)
		return
	}
	imageType, err := models.GetImageType(requestImageId, config.Prog.Db)
	if err != nil {
		// will be a wrong image id
		http.Error(w, "Unable to process the request", http.StatusBadRequest)
		return
	}
	imageName := fmt.Sprintf("%s.%s", requestImageId, imageType)
	detail := fmt.Sprintf("Image URL: http://localhost:8080/%s", imageName)
	apiResponse := createResponse(detail, "Success", http.StatusOK)
	jsonResponseDecorator(apiResponse, w)
}
