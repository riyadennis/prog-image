package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/prog-image/middleware"
	"github.com/prog-image/models"
)

func GetImageHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	imageId := params.ByName("image_id")
	config, err := middleware.GetConfigFromContext(req.Context())
	if err != nil {
		http.Error(w, "Unable to process the request", http.StatusInternalServerError)
		return
	}
	imageType, err := models.GetImageType(imageId, config.Prog.Db)
	if err != nil {
		// will be a wrong image id
		http.Error(w, "Unable to process the request", http.StatusBadRequest)
		return
	}
	imageName := fmt.Sprintf("%s.%s", imageId, imageType)
	detail := fmt.Sprintf("Image URL: http://localhost:8080/%s", imageName)
	apiResponse := createResponse(detail, "Success", http.StatusOK)
	jsonResponseDecorator(apiResponse, w)
}
