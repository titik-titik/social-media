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

func DeserializePostRows(rows *sql.Rows) []*entity.Post {
	var foundPosts []*entity.Post
	for rows.Next() {
		foundPost := &entity.Post{}
		scanErr := rows.Scan(
			&foundPost.Id,
			&foundPost.UserId,
			&foundPost.Description,
			&foundPost.ImageUrl,
			&foundPost.UpdatedAt,
			&foundPost.CreatedAt,
			&foundPost.DeletedAt,
		)
		if scanErr != nil {
			panic(scanErr)
		}
		foundPosts = append(foundPosts, foundPost)
	}
	return foundPosts
}

func (searchRepository *SearchRepository) FindAllUser(begin *sql.Tx) []*entity.User {
	rows, queryErr := begin.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user", nil,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundAllUser := DeserializeUserRows(rows)
	return foundAllUser

}

func (searchRepository *SearchRepository) FindAllPostByUserId(begin *sql.Tx, id string) []*entity.Post {
	rows, queryErr := begin.Query(
		"SELECT id, user_id, description, image_url, created_at, updated_at, deleted_at FROM \"post\" where user_id = ? LIMIT 1", id,
	)
	if queryErr != nil {
		panic(queryErr)
	}
	foundAllPosts := DeserializePostRows(rows)
	return foundAllPosts
}
