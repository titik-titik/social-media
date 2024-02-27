package model

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
