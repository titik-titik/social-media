package response

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
	Code    int    `json:"code"`
}

func NewResponse[T any](message string, data T, code int) *Response[T] {
	response := &Response[T]{
		Message: message,
		Data:    data,
		Code:    code,
	}
	return response
}

func ErrorResponse(w http.ResponseWriter, message string, code int) {
	resp := NewResponse(message, nil, code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

func SuccessResponse(w http.ResponseWriter, message string, data interface{}, code int) {
	resp := NewResponse(message, data, code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}
