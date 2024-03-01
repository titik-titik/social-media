package converter

import (
	"social-media/internal/entity"
	"social-media/internal/model/response"
)

func PostToResponse(post *entity.Post) *response.PostResponse {
	return &response.PostResponse{
		UserId:      post.UserId,
		ImageUrl:    post.ImageUrl,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}
