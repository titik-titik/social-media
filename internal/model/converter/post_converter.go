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

func PostToResponses(post []entity.Post) *[]response.PostResponse {
	posts := make([]response.PostResponse, 0, len(post))

	for i := range posts {
		posts = append(posts, *PostToResponse(&post[i]))
	}

	return &posts
}
