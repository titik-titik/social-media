package seeder

import (
	"social-media/internal/config"
	"social-media/test/mock"
)

type PostSeeder struct {
	DatabaseConfig *config.DatabaseConfig
	PostMock       *mock.PostMock
}

func NewPostSeeder(
	databaseConfig *config.DatabaseConfig,
	userSeeder *UserSeeder,
) *PostSeeder {
	postSeeder := &PostSeeder{
		DatabaseConfig: databaseConfig,
		PostMock:       mock.NewPostMock(userSeeder.UserMock),
	}
	return postSeeder
}

func (postSeeder *PostSeeder) Up() {
	for _, post := range postSeeder.PostMock.Data {
		begin, beginErr := postSeeder.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if beginErr != nil {
			panic(beginErr)
		}

		_, err := begin.Query(
			"INSERT INTO \"post\" (id, user_id, image_url, description, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7);",
			post.Id,
			post.UserId,
			post.ImageUrl,
			post.Description,
			post.CreatedAt,
			post.UpdatedAt,
			post.DeletedAt,
		)
		if err != nil {
			panic(err)
		}

		commitErr := begin.Commit()
		if commitErr != nil {
			panic(commitErr)
		}
	}
}

func (postSeeder *PostSeeder) Down() {
	for _, post := range postSeeder.PostMock.Data {
		begin, beginErr := postSeeder.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if beginErr != nil {
			panic(beginErr)
		}

		_, err := begin.Query(
			"DELETE FROM \"post\" WHERE id = $1 LIMIT 1;",
			post.Id,
		)
		if err != nil {
			panic(err)
		}

		commitErr := begin.Commit()
		if commitErr != nil {
			panic(commitErr)
		}
	}
}
