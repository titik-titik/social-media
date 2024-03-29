package repository

import (
	"database/sql"
	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"social-media/internal/entity"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	userRepository := &UserRepository{}
	return userRepository
}

func DeserializeUserRows(rows *sql.Rows) []*entity.User {
	var foundUsers []*entity.User
	for rows.Next() {
		foundUser := &entity.User{}
		scanErr := rows.Scan(
			&foundUser.Id,
			&foundUser.Name,
			&foundUser.Username,
			&foundUser.Email,
			&foundUser.Password,
			&foundUser.AvatarUrl,
			&foundUser.Bio,
			&foundUser.IsVerified,
			&foundUser.CreatedAt,
			&foundUser.UpdatedAt,
			&foundUser.DeletedAt,
		)
		if scanErr != nil {
			panic(scanErr)
		}
		foundUsers = append(foundUsers, foundUser)
	}
	return foundUsers
}

func (userRepository *UserRepository) FindOneById(tx *sql.Tx, id string) (result *entity.User, err error) {
	var rows *sql.Rows
	var queryErr error
	_ = crdb.Execute(func() error {
		rows, queryErr = tx.Query(
			`SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM "user" WHERE id=$1 LIMIT 1;`,
			id,
		)
		return queryErr
	})
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		result = nil
		err = nil
		return result, err
	}

	result = foundUsers[0]
	err = nil
	return result, err
}

func (userRepository *UserRepository) FindOneByUsername(tx *sql.Tx, username string) (result *entity.User, err error) {
	rows, queryErr := tx.Query(
		`SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM "user" WHERE username=$1 LIMIT 1;`,
		username,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		result = nil
		err = nil
		return result, err
	}

	result = foundUsers[0]
	err = nil
	return result, err
}

func (userRepository *UserRepository) FindOneByEmail(tx *sql.Tx, email string) (result *entity.User, err error) {
	rows, queryErr := tx.Query(
		`SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM "user" WHERE email=$1 LIMIT 1;`,
		email,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		result = nil
		err = nil
		return result, err
	}

	result = foundUsers[0]
	err = nil
	return result, err
}

func (userRepository *UserRepository) FindOneByEmailAndPassword(tx *sql.Tx, email string, password string) (result *entity.User, err error) {
	rows, queryErr := tx.Query(
		`SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM "user" WHERE email=$1 AND password=$2 LIMIT 1;`,
		email,
		password,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		result = nil
		err = nil
		return result, err
	}

	result = foundUsers[0]
	err = nil
	return result, err
}

func (userRepository *UserRepository) FindOneByUsernameAndPassword(tx *sql.Tx, username string, password string) (result *entity.User, err error) {
	rows, queryErr := tx.Query(
		`SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM "user" WHERE username=$1 AND password=$2 LIMIT 1;`,
		username,
		password,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		result = nil
		err = nil
		return result, err
	}

	result = foundUsers[0]
	err = nil
	return result, err
}

func (userRepository *UserRepository) CreateOne(tx *sql.Tx, toCreateUser *entity.User) (result *entity.User, err error) {
	_, queryErr := tx.Query(
		`INSERT INTO "user" (id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`,
		toCreateUser.Id,
		toCreateUser.Name,
		toCreateUser.Username,
		toCreateUser.Email,
		toCreateUser.Password,
		toCreateUser.AvatarUrl,
		toCreateUser.Bio,
		toCreateUser.IsVerified,
		toCreateUser.CreatedAt,
		toCreateUser.UpdatedAt,
		toCreateUser.DeletedAt,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	result = toCreateUser
	err = nil
	return result, err
}

func (userRepository *UserRepository) PatchOneById(tx *sql.Tx, id string, toPatchUser *entity.User) (result *entity.User, err error) {
	_, queryErr := tx.Query(
		`UPDATE "user" SET id=$1, name=$2, username=$3, email=$4, password=$5, avatar_url=$6, bio=$7, is_verified=$8, created_at=$9, updated_at=$10, deleted_at=$11 WHERE id = $12 LIMIT 1;`,
		toPatchUser.Id,
		toPatchUser.Name,
		toPatchUser.Username,
		toPatchUser.Email,
		toPatchUser.Password,
		toPatchUser.AvatarUrl,
		toPatchUser.Bio,
		toPatchUser.IsVerified,
		toPatchUser.CreatedAt,
		toPatchUser.UpdatedAt,
		toPatchUser.DeletedAt,
		id,
	)

	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	result = toPatchUser
	err = nil
	return result, err
}

func (userRepository *UserRepository) DeleteOneById(tx *sql.Tx, id string) (result *entity.User, err error) {
	rows, queryErr := tx.Query(
		`DELETE FROM "user" WHERE id=$1 LIMIT 1 RETURNING id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at`,
		id,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		result = nil
		err = nil
		return result, err
	}

	result = foundUsers[0]
	err = nil
	return result, err
}
