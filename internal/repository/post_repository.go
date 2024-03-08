package repository

import (
	"context"
	"errors"
	"github.com/guregu/null"
	"github.com/jackc/pgx/v5"
	"social-media/internal/entity"
	"time"
)

type PostRepository struct {
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func DeserializePostRows(rows pgx.Rows) []*entity.Post {
	foundPosts, collectRowErr := pgx.CollectRows(rows, pgx.RowToStructByName[*entity.Post])
	if collectRowErr != nil {
		panic(collectRowErr)
	}
	return foundPosts

}

func (p *PostRepository) Create(db pgx.Tx, post *entity.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Query(ctx, "INSERT INTO post (id,user_id,image_url,description,created_at,updated_at,deleted_at) VALUES ($1,$2,$3,$4,$5,$6,$7)", post.Id, post.UserId, post.ImageUrl, post.Description, post.CreatedAt, post.UpdatedAt, post.DeletedAt)

	if err != nil {
		return errors.New("failed to create new post")
	}

	return nil
}

func (p *PostRepository) Get(db pgx.Tx, post *entity.Post, postId null.String) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx, "SELECT * FROM post WHERE id=$1", postId.String).Scan(&post.Id, &post.ImageUrl, &post.UserId, &post.Description, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
