package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"io/ioutil"
	"github.com/prog-image/service"
	"github.com/satori/go.uuid"
	"fmt"
)

type UploadedImage struct {
	Uri string `json:"uri"`
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
	////need to use this to store the image in db also
	filename := fmt.Sprintf("%s.jpg", uuid.Must(uuid.NewV1(), nil))
	image := service.NewImage("../images", filename, uploadedImage.Uri)
	created, err := image.CreateImage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if created == false {
		w.Write([]byte("Unable to save the image"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//
	//w.Write([]byte("Image saved successfully"))
}
