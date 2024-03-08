package seeder

import (
	"fmt"
)

type AllSeeder struct {
	User *UserSeeder
	Post *PostSeeder
}

func NewAllSeeder(
	user *UserSeeder,
	post *PostSeeder,
) *AllSeeder {
	allSeeder := &AllSeeder{
		User: user,
		Post: post,
	}
	return allSeeder
}

func (allSeeder *AllSeeder) Up() {
	fmt.Println("Seeder up started.")
	allSeeder.User.Up()
	allSeeder.Post.Up()
	fmt.Println("Seeder up finished.")
}

func (allSeeder *AllSeeder) Down() {
	fmt.Println("Seeder down started.")
	allSeeder.User.Down()
	allSeeder.Post.Down()
	fmt.Println("Seeder down finished.")
}
