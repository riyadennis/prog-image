package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	state := struct {
		Status string    `json:"status"`
		Errors [0]string `json:"errors"`
	}{Status: "OK"}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(state); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
