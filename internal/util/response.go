package util

import (
	"encoding/json"
	"net/http"

	"github.com/shironxn/go-clean-arch/internal/domain"
)

func ResponseError(w http.ResponseWriter, statusCode int, message string) {
	ResponseJSON(w, statusCode, domain.ErrorResponse{
		Error: message,
	})
}

func ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
