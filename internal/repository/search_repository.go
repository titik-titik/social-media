package repository

import (
	"database/sql"
	// "social-media/internal/config"
	"social-media/internal/config"
	"social-media/internal/entity"
)

type SearchRepository struct {
	Database *config.DatabaseConfig
}

func NewSearchRepository(database *config.DatabaseConfig) *SearchRepository {
	searchRepository := &SearchRepository{
		Database: database,
	}
	return searchRepository
}

func deserializeRowsForPost(rows *sql.Rows) []*entity.Post {
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

		foundPost.CreatedAt.Time = foundPost.CreatedAt.Time.UTC()
		foundPost.UpdatedAt.Time = foundPost.UpdatedAt.Time.UTC()
		foundPost.DeletedAt.Time = foundPost.DeletedAt.Time.UTC()
		foundPosts = append(foundPosts, foundPost)
	}
	return foundPosts
}

func (searchRepository *SearchRepository) FindAllUser() []*entity.User {
	begin, beginErr := searchRepository.Database.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}
	rows, queryErr := begin.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user", nil,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundAllUser := deserializeRows(rows)
	return foundAllUser

}

func (searchRepository *SearchRepository) FindAllPostByUserId(id string) []*entity.Post {
	begin, beginErr := searchRepository.Database.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}
	rows, queryErr := begin.Query(
		"SELECT id, user_id, description, image, created_at, updated_at, deleted_at FROM post where user_id = ? LIMIT 1", id,
	)
	if queryErr != nil {
		panic(queryErr)
	}
	foundAllPosts := deserializeRowsForPost(rows)
	return foundAllPosts
}
