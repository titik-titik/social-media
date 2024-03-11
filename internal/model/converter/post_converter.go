package converter

import (
	"social-media/internal/entity"
	"social-media/internal/model/response"
)

func PostToResponse(post *entity.Post) *response.PostResponse {
	return &response.PostResponse{
		ID:          post.Id,
		UserId:      post.UserId,
		ImageUrl:    post.ImageUrl,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}

func PostToResponses(posts []*entity.Post) []*response.PostResponse {
	var newPosts []*response.PostResponse

	for _, p := range posts {
		newPosts = append(newPosts, PostToResponse(p))
	}

	return newPosts
}
