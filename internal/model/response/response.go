package response

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Code    int    `json:"code"`
}

func NewResponse[T any](w http.ResponseWriter, message string, data T, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	response := &Response[T]{
		Message: message,
		Data:    data,
		Code:    statusCode,
	}

	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
