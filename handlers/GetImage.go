package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func GetImageHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	imageId := params.ByName("image_id")
	fmt.Fprint(w, imageId)
}