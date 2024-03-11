package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/guregu/null"
	"github.com/rs/zerolog"
	"social-media/internal/entity"
	"strings"
)

type PostRepository struct {
	Log *zerolog.Logger
}

func NewPostRepository(log *zerolog.Logger) *PostRepository {
	return &PostRepository{
		Log: log,
	}
}

func (p *PostRepository) Create(db *sql.Tx, post *entity.Post) error {
	_, err := db.Query("INSERT INTO post (id,user_id,image_url,description,created_at,updated_at,deleted_at) VALUES ($1,$2,$3,$4,$5,$6,$7)", post.Id, post.UserId, post.ImageUrl, post.Description, post.CreatedAt, post.UpdatedAt, post.DeletedAt)

	if err != nil {
		return errors.New("failed to create new post")
	}

	return nil
}

func (p *PostRepository) FindByID(db *sql.Tx, post *entity.Post, postId null.String) error {
	if err := db.QueryRow("SELECT id,user_id,image_url,description,created_at,updated_at,deleted_at FROM post WHERE id=$1", postId.String).Scan(&post.Id, &post.UserId, &post.ImageUrl, &post.Description, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func DeserializePostRows(rows *sql.Rows, foundPosts *[]entity.Post) error {
	for rows.Next() {
		foundPost := new(entity.Post)
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
			return scanErr
		}
		*foundPosts = append(*foundPosts, *foundPost)
	}
	return nil
}

func (p *PostRepository) Get(db *sql.Tx, posts *[]entity.Post, order string, limit int8, offset int64) error {
	query := fmt.Sprintf(`SELECT id,user_id,image_url,description,created_at,updated_at,deleted_at FROM post ORDER BY updated_at %s LIMIT %d OFFSET %d`, strings.ToUpper(order), limit, offset)

	rows, err := db.Query(query)

	if err != nil {
		return err
	}

	if err = DeserializePostRows(rows, posts); err != nil {
		return err
	}

	return nil
}

func (p PostRepository) Update(db *sql.Tx, posts *entity.Post, postID string) error {
	_, err := db.Query("UPDATE post set image_url=$1, description=$2,updated_at=$3 WHERE id = $4", posts.ImageUrl, posts.Description, posts.UpdatedAt, postID)

	if err != nil {
		return err
	}

	return nil
}

func (p PostRepository) Delete(db *sql.Tx, postID string) error {
	_, err := db.Query("DELETE FROM post WHERE id = $1", postID)

	if err != nil {
		return err
	}

	return nil
}

func (p PostRepository) CountByID(db *sql.Tx, postID string) (int64, error) {
	var total int64

	if err := db.QueryRow("SELECT COUNT('id') FROM post WHERE id = $1", postID).Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
