package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/prog-image/middleware"
	"github.com/prog-image/models"
	"github.com/prog-image/service"
	"github.com/pkg/errors"
	"os"
	"io"
	"image/png"
	"github.com/sirupsen/logrus"
	"image"
	_ "image/jpeg"
)

func GetImageHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestImageId := params.ByName("image_id")
	requestImageType := req.URL.Query().Get("type")
	config, err := middleware.GetConfigFromContext(req.Context())
	if err != nil {
		logrus.Errorf("Unable to fetch config got error %s", err.Error())
		http.Error(w, "Unable to process the request", http.StatusInternalServerError)
		return
	}
	imageType, err := models.GetImageType(requestImageId, config.Prog.Db)
	if err != nil {
		// will be a wrong image id
		logrus.Errorf("Image type not found %s", err.Error())
		http.Error(w, "Unable to process the request", http.StatusBadRequest)
		return
	}
	imageName := fmt.Sprintf("%s.%s", requestImageId, imageType)
	//detail for get requests with out conversion
	detail := fmt.Sprintf("Image URL: http://localhost:8080/%s", imageName)
	if requestImageType != "" {
		err = CreateNewImageForImageType(requestImageId,requestImageType, imageType, config)
		if err != nil {
			// wrong image format and not able to create new image
			http.Error(w, "Unable to process the request", http.StatusBadRequest)
			return
		}
		// detail for get request with conversion
		detail = fmt.Sprintf("Image URL: http://localhost:8080/%s.%s", requestImageId, requestImageType)
	}

	apiResponse := createResponse(detail, "Success", http.StatusOK)
	jsonResponseDecorator(apiResponse, w)
}
func CreateNewImageForImageType(requestImageId, requestImageType, existingImageType string, config *middleware.Config) (error) {
	if !service.ValidateFileType(requestImageType, config) {
		return errors.New("invalid image type")
	}
	var oldFile io.Reader
	newFilePath := fmt.Sprintf("%s/%s.%s", config.Prog.Folder, requestImageId, requestImageType)
	oldFilePath := fmt.Sprintf("%s/%s.%s", config.Prog.Folder, requestImageId, existingImageType)
	oldFile, err := os.Open(oldFilePath)
	if err != nil {
		return  err
	}
	newFile, err := os.Create(newFilePath)
	if err != nil {
		return err
	}
	if requestImageType == "png" {
		err := convertToPNG(newFile, oldFile)
		if err != nil {
			return err
		}
	}

	return nil
}
func convertToPNG(newFile io.Writer, oldFile io.Reader) error {
	img, _, err := image.Decode(oldFile)
	if err != nil {
		return err
	}
	return png.Encode(newFile, img)
}
