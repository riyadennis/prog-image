package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"io/ioutil"
	. "github.com/prog-image/service"
)

type UploadedImage struct {
	Uri string `json:"uri"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uploaded, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	uploadedImage := UploadedImage{}
	err = json.Unmarshal(uploaded, &uploadedImage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	image := Image{Path: "images/", Type:"jpg", Source : uploadedImage.Uri}
	image.CreateImage()
	w.Write([]byte("hello"))
}
