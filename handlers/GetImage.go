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
	"image/jpeg"
	"io"
	"image/png"
	"github.com/sirupsen/logrus"
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
		http.Error(w, "Unable to process the request", http.StatusBadRequest)
		return
	}
	imageName := fmt.Sprintf("%s.%s", requestImageId, imageType)
	detail := fmt.Sprintf("Image URL: http://localhost:8080/%s", imageName)
	if requestImageType != "" {
		err = CreateNewImageForImageType(requestImageId,requestImageType, imageType, config)
		if err != nil {
			// wrong image format and not able to create new image
			http.Error(w, "Unable to process the request", http.StatusBadRequest)
			return
		}
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
	if requestImageType == "png" && existingImageType == "jpg" {
		err := convertJPGTOPNG(oldFile, newFile)
		if err != nil {
			return err
		}
	}

	return nil
}
func convertJPGTOPNG(oldFile io.Reader, newFile io.Writer) (error) {
	img, err := jpeg.Decode(oldFile)
	if err != nil {
		return err
	}
	err = png.Encode(newFile, img)
	if err != nil {
		return err
	}
	return nil
}
