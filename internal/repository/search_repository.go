package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"

	"social-media/internal/entity"
)

type SearchRepository struct {
}

func NewSearchRepository() *SearchRepository {
	searchRepository := &SearchRepository{}
	return searchRepository
}

func (searchRepository *SearchRepository) FindAllUser(begin pgx.Tx) []*entity.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user", nil,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundAllUser := DeserializeUserRows(rows)
	return foundAllUser

}

func (searchRepository *SearchRepository) FindAllPostByUserId(begin pgx.Tx, id string) []*entity.Post {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"SELECT id, user_id, description, image_url, created_at, updated_at, deleted_at FROM \"post\" where user_id = ? LIMIT 1", id,
	)

	if queryErr != nil {
		panic(queryErr)
	}
	foundAllPosts := DeserializePostRows(rows)
	return foundAllPosts
}
