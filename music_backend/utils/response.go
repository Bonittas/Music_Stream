package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := ErrorResponse{Message: message}
	response, _ := json.Marshal(errorResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}