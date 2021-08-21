package utils

import (
	"books-list/err"
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, status int, err err.Error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

type ServerMessage struct {
	Message string `json:"message"`
}
