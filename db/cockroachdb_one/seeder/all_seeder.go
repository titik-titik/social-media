package seeder

import (
	"fmt"
)

type AllSeeder struct {
	UserSeeder *UserSeeder
	AuthSeeder *AuthSeeder
}

func NewAllSeeder(
	userSeeder *UserSeeder,
	AuthSeeder *AuthSeeder,
) *AllSeeder {
	allSeeder := &AllSeeder{
		UserSeeder: userSeeder,
		AuthSeeder: AuthSeeder,
	}
	return allSeeder
}

func (s *AllSeeder) Up() {
	fmt.Println("Seeder up started.")
	s.UserSeeder.Up()
	s.AuthSeeder.Up()
	fmt.Println("Seeder up finished.")
}

func (s *AllSeeder) Down() {
	fmt.Println("Seeder down started.")
	s.UserSeeder.Down()
	s.AuthSeeder.Down()
	fmt.Println("Seeder down finished.")
}
