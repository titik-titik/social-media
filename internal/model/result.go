package model

type Result[T any] struct {
	Code    int
	Message string
	Data    T
}

func NewResult[T any](code int, message string, data T) *Result[T] {
	result := &Result[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return result
}
