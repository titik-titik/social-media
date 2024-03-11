package request

type GetAllUserRequest struct {
	Limit  string `json:"limit" validate:"required"`
	Offset string `json:"offset" validate:"gte=0"`
	Order  string `json:"order" validate:"oneof=desc asc"`
}
