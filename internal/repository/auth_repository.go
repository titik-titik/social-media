package repository

import (
	"database/sql"
	"time"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
func (ar *AuthRepository) Register(id, name, email, hashedPassword string, createdAt, updatedAt time.Time) error {
	createSQL := `
	    INSERT INTO users (id, name, email, password, created_at, updated_at)
	    VALUES (?, ?, ?, ?, CONVERT_TZ(?, '+00:00', '+07:00'), CONVERT_TZ(?, '+00:00', '+07:00'))
	`

	_, err := ar.db.Exec(createSQL, id, name, email, hashedPassword, createdAt.UTC(), updatedAt.UTC())
	return err
}
