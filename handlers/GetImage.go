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
	"image/jpeg"
	"image/gif"
)

type FileConverter struct {
	NewFileType string
	OldFileType string
	NewFile     io.Writer
	OldFile     io.Reader
}

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
	//detail for get requests with out conversion
	detail := fmt.Sprintf("Image URL: %s",GetLocalImageURL(config, requestImageId, imageType))
	if requestImageType != "" {
		err = CreateNewImageForImageType(requestImageId, requestImageType, imageType, config)
		if err != nil {
			// wrong image format and not able to create new image
			http.Error(w, "Unable to process the request", http.StatusBadRequest)
			return
		}
		// detail for get request with conversion
		detail = fmt.Sprintf("Image URL: %s",GetLocalImageURL(config, requestImageId, requestImageType))
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
		return err
	}
	newFile, err := os.Create(newFilePath)
	if err != nil {
		return err
	}
	fileConverter := FileConverter{
		NewFileType: requestImageType,
		NewFile:newFile,
		OldFileType: existingImageType,
		OldFile: oldFile,
	}
	err = fileConverter.ConvertFileFromOnTypeToAnother()
	if err != nil {
		return err
	}
	err = models.SaveImageType(requestImageId, requestImageType, config.Prog.Db)
	if err != nil {
		return err
	}
	return nil
}
func (fc FileConverter) ConvertFileFromOnTypeToAnother() error {
	img, _, err := image.Decode(fc.OldFile)
	if err != nil {
		return err
	}
	switch fc.NewFileType {
	case "jpg":
		return jpeg.Encode(fc.NewFile, img, nil)
	case "png" :
		return png.Encode(fc.NewFile, img)
	case "gif":
		return gif.Encode(fc.NewFile, img, nil)
	default:
		return errors.New("Invalid file format")
	}
	return nil
}