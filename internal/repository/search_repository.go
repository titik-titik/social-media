package repository

import (
	"database/sql"
	"fmt"

	"social-media/internal/entity"
)

type SearchRepository struct {
}

func NewSearchRepository() *SearchRepository {
	searchRepository := &SearchRepository{}
	return searchRepository
}

func (searchRepository *SearchRepository) FindManyUser(begin *sql.Tx) (result []*entity.User, err error) {
	rows, queryErr := begin.Query(
		`SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM "user";`,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	foundAllUser := DeserializeUserRows(rows)

	result = foundAllUser
	err = nil
	fmt.Println(foundAllUser)
	return result, err
}

func (searchRepository *SearchRepository) FindManyPostByUserId(begin *sql.Tx, id string) (result []*entity.Post, err error) {
	rows, queryErr := begin.Query(
		`SELECT id, user_id, image_url, description, created_at, updated_at, deleted_at FROM "post" where user_id = $1`, id,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}
	var posts []*entity.Post
	deserializerErr := DeserializePostRows(rows, &posts)
	if deserializerErr != nil {
		result = nil
		err = deserializerErr
		return result, err
	}

	result = posts
	err = nil
	return
}

func (searchRepository *SearchRepository) FindPostByDescription(begin *sql.Tx, description string) (result []*entity.Post, err error) {
	rows, queryErr := begin.Query(
		`SELECT id, user_id, description, image_url, created_at, updated_at, deleted_at FROM "post" where description = $1`, description,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}

	var posts []*entity.Post
	deserializerErr := DeserializePostRows(rows, &posts)
	if deserializerErr != nil {
		result = nil
		err = deserializerErr
		return result, err
	}

	result = posts
	err = nil
	return result, err
}
