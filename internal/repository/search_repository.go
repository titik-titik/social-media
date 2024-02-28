package repository

import (
	"database/sql"
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

func rowMapperForPost(rows *sql.Rows) *entity.Post {
	foundUser := &entity.Post{}
	scanErr := rows.Scan(
		&foundUser.Id,
		&foundUser.User_Id,
		&foundUser.Description,
		&foundUser.Image,
		&foundUser.CreatedAt,
		&foundUser.UpdatedAt,
		&foundUser.DeletedAt,
	)
	if scanErr != nil {
		panic(scanErr)
	}
	return foundUser
}

func (searchRepository *SearchRepository) FindAllUser() *[]entity.User {
	rows, queryErr := searchRepository.Database.MariaDbOneDatabase.Db.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user", nil,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	var foundAllUser []entity.User
	for rows.Next() {
		foundUser := rowMapper(rows)
		foundAllUser = append(foundAllUser, *foundUser)
	}
	return &foundAllUser

}

func (searchRepository *SearchRepository) FindAllPostByUserId(id string) *[]entity.Post {
	rows, queryErr := searchRepository.Database.MariaDbOneDatabase.Db.Query(
		"SELECT id, user_id, description, image, created_at, updated_at, deleted_at FROM post where user_id = ? LIMIT 1", id,
	)
	if queryErr != nil {
		panic(queryErr)
	}
	var foundAllPost []entity.Post = nil
	for rows.Next() {
		foundPost := rowMapperForPost(rows)
		foundAllPost = append(foundAllPost, *foundPost)
	}
	return &foundAllPost
}
