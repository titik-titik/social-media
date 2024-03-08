package seeder

import (
	"context"
	"social-media/internal/config"
	"social-media/test/mock"
	"time"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, post := range postSeeder.PostMock.Data {
		connection, acquireErr := postSeeder.DatabaseConfig.CockroachDatabase.Pool.Acquire(ctx)
		if acquireErr != nil {
			panic(acquireErr)
		}
		begin, beginErr := connection.Begin(ctx)
		if beginErr != nil {
			panic(beginErr)
		}

		_, err := begin.Query(
			ctx,
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

		commitErr := begin.Commit(ctx)
		if commitErr != nil {
			panic(commitErr)
		}

		connection.Release()
	}
}

func (postSeeder *PostSeeder) Down() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, post := range postSeeder.PostMock.Data {
		connection, acquireErr := postSeeder.DatabaseConfig.CockroachDatabase.Pool.Acquire(ctx)
		if acquireErr != nil {
			panic(acquireErr)
		}
		begin, beginErr := connection.Begin(ctx)
		if beginErr != nil {
			panic(beginErr)
		}

		_, queryErr := begin.Query(
			ctx,
			"DELETE FROM \"post\" WHERE id = $1 LIMIT 1;",
			post.Id,
		)
		if queryErr != nil {
			panic(queryErr)
		}

		commitErr := begin.Commit(ctx)
		if commitErr != nil {
			panic(commitErr)
		}

		connection.Release()
	}
}
