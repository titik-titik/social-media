package entity

import "time"

// User is a struct that represents a user entity
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	Bio       string    `json:"bio"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
