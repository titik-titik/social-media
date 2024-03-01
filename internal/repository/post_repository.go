package repository

import (
	"database/sql"
	"social-media/internal/config"
	"social-media/internal/entity"
)

type PostRepository struct {
	Database *config.DatabaseConfig
}

func NewPostRepository(database *config.DatabaseConfig) *PostRepository {
	return &PostRepository{
		Database: database,
	}
}

func (p *PostRepository) Create(db *sql.DB, post *entity.Post) error {
	_, err := db.Query("INSERT INTO post (id,user_id,image_url,description,created_at,updated_at,deleted_at) VALUES ($1,$2,$3,$4,$5,$6,$7)", post.Id, post.UserId, post.ImageUrl, post.Description, post.CreatedAt, post.UpdatedAt, post.DeletedAt)

	if err != nil {
		panic("Failed to create new post")
	}

	return nil
}
