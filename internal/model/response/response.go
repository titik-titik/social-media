package response

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Message string      `json:"message,omitempty"`
	Data    T           `json:"data,omitempty"`
	Code    int         `json:"code,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func NewResponse[T any](w http.ResponseWriter, response *Response[T]) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(response.Code)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
