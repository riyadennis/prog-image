package handlers

import (
	"net/http"
	"os"
	"github.com/sirupsen/logrus"
	"io"
	"fmt"
)
type FileConf struct{
	path  string
}
func ( fc *FileConf) FileServer(w http.ResponseWriter, r *http.Request) {
	fullPath := fmt.Sprintf("%s/%s", fc.path, r.URL.Path)
	f, err := os.Open(fullPath)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Invalid file name"))
		logrus.Errorf("Unable to open images got error %s", err.Error())
	}
	w.Header().Add("Content-Type", "image/jpeg")
	io.Copy(w, f)
}
