package response

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewResponse[T any](message string, data T) *Response[T] {
	response := &Response[T]{
		Message: message,
		Data:    data,
	}
	return response
}

// ErrorResponse returns a JSON response with an error message.
func ErrorResponse(w http.ResponseWriter, message string, status int) {
	type Response struct {
		Error   bool   `json:"Error"`
		Message string `json:"message"`
	}

	response := Response{
		Error:   true, // This should probably be true to indicate an error
		Message: message,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		// Recursive call could potentially cause a stack overflow if json.Marshal fails repeatedly
		// Consider logging the error instead and return a simple error message or status code
		ErrorResponse(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Corrected "internal/json" to "application/json"
	w.WriteHeader(status)
	w.Write(responseJSON)
}

func SuccessResponse(w http.ResponseWriter, message string, data interface{}, status int) {
	type Response struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		ErrorResponse(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Corrected "internal/json" to "application/json"
	w.WriteHeader(status)
	w.Write(responseJSON)
}

func OtherResponses(w http.ResponseWriter, message string, status int) {
	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	response := Response{
		Success: true,
		Message: message,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		ErrorResponse(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Corrected "internal/json" to "application/json"
	w.WriteHeader(status)
	w.Write(responseJSON)
}
