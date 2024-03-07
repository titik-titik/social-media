package model

type TokenResult struct {
	Code    int
	Message string
	Data    string
}

func NewTokenResult(code int, message string, token string) *TokenResult {
	result := &TokenResult{
		Code:    code,
		Message: message,
		Data:    token,
	}
	return result
}
