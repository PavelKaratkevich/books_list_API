package utils

import (
	"books-list/domain"
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, status int, err domain.Error) {
	w.Header().Add("Content-Type", "routes/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "routes/json")
	json.NewEncoder(w).Encode(data)
}
