package repository

import (
	"database/sql"
	"errors"
	"github.com/guregu/null"
	"social-media/internal/entity"
)

type PostRepository struct {
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func DeserializePostRows(rows *sql.Rows) []*entity.Post {
	var foundPosts []*entity.Post
	for rows.Next() {
		foundPost := &entity.Post{}
		scanErr := rows.Scan(
			&foundPost.Id,
			&foundPost.UserId,
			&foundPost.ImageUrl,
			&foundPost.Description,
			&foundPost.CreatedAt,
			&foundPost.UpdatedAt,
			&foundPost.DeletedAt,
		)
		if scanErr != nil {
			panic(scanErr)
		}
		foundPosts = append(foundPosts, foundPost)
	}
	return foundPosts

}

func (p *PostRepository) Create(db *sql.Tx, post *entity.Post) error {
	_, queryErr := db.Query("INSERT INTO \"post\" (id,user_id,image_url,description,created_at,updated_at,deleted_at) VALUES ($1,$2,$3,$4,$5,$6,$7)", post.Id, post.UserId, post.ImageUrl, post.Description, post.CreatedAt, post.UpdatedAt, post.DeletedAt)
	if queryErr != nil {
		return errors.New("failed to create new post")
	}

	return nil
}

func (p *PostRepository) Get(db *sql.Tx, post *entity.Post, postId null.String) error {
	queryErr := db.QueryRow("SELECT * FROM \"post\" WHERE id=$1", postId.String).Scan(&post.Id, &post.ImageUrl, &post.UserId, &post.Description, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
	if queryErr != nil {
		return errors.New("failed to get post")
	}

	return nil
}
