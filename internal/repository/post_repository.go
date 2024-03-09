package repository

import (
	"database/sql"
	"errors"
	"github.com/guregu/null"
	"social-media/internal/entity"
	"social-media/tool"
)

type PostRepository struct {
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (p *PostRepository) Create(db *sql.Tx, post *entity.Post) error {
	_, err := db.Query("INSERT INTO post (id,user_id,image_url,description,created_at,updated_at,deleted_at) VALUES ($1,$2,$3,$4,$5,$6,$7)", post.Id, post.UserId, post.ImageUrl, post.Description, post.CreatedAt, post.UpdatedAt, post.DeletedAt)

	if err != nil {
		return errors.New("failed to create new post")
	}

	return nil
}

func (p *PostRepository) FindByID(db *sql.Tx, post *entity.Post, postId null.String) error {
	if err := db.QueryRow("SELECT * FROM post WHERE id=$1", postId.String).Scan(&post.Id, &post.ImageUrl, &post.UserId, &post.Description, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (p *PostRepository) Get(db *sql.Tx, posts *[]entity.Post, order string, limit int8, offset int64) error {
	rows, err := db.Query("SELECT id,user_id,image_url,description,created_at,updated_at FROM post ORDER BY $1 limit $2 offset $3", order, limit, offset)

	if err != nil {
		return err
	}

	posts = tool.DeserializeRows(rows, &entity.Post{}).(*[]entity.Post)

	return nil
}
