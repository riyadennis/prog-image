package handlers

import (
	"net/http"
	"os"
	"github.com/sirupsen/logrus"
	"io"
	"fmt"
	"strings"
)
type FileConf struct{
	path  string
}
func ( fc *FileConf) FileServer(w http.ResponseWriter, r *http.Request) {
	//just initailise it in case
	fileType := "jpg"
	fileTYpeFromRequest := strings.Split(r.URL.Path, ".")
	if fileTYpeFromRequest[1] != ""{
		fileType = fileTYpeFromRequest[1]
	}

	fullPath := fmt.Sprintf("%s/%s", fc.path, r.URL.Path)
	f, err := os.Open(fullPath)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Invalid file name"))
		logrus.Errorf("Unable to open images got error %s", err.Error())
	}

	imageType := fmt.Sprintf("image/%s", fileType)
	w.Header().Add("Content-Type", imageType)
	io.Copy(w, f)
}
