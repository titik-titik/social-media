package repository

import (
	"database/sql"
	"social-media/internal/entity"
)

type SearchRepository struct {
}

func NewSearchRepository() *SearchRepository {
	searchRepository := &SearchRepository{}
	return searchRepository
}

func (searchRepository *SearchRepository) FindAllUser(begin *sql.Tx) (result []*entity.User, err error) {
	rows, queryErr := begin.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user", nil,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	foundAllUser := DeserializeUserRows(rows)

	result = foundAllUser
	err = nil
	return result, err
}

func (searchRepository *SearchRepository) FindAllPostByUserId(begin *sql.Tx, id string) (result []*entity.Post, err error) {
	rows, queryErr := begin.Query(
		"SELECT id, user_id, description, image_url, created_at, updated_at, deleted_at FROM \"post\" where user_id = ? LIMIT 1", id,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	foundAllPosts := DeserializePostRows(rows)

	result = foundAllPosts
	err = nil
	return result, err
}
