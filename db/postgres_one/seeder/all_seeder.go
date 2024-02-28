package seeder

import (
	"fmt"
)

type AllSeeder struct {
	UserSeeder *UserSeeder
}

func NewAllSeeder(
	userSeeder *UserSeeder,
) *AllSeeder {
	allSeeder := &AllSeeder{
		UserSeeder: userSeeder,
	}
	return allSeeder
}

func (s *AllSeeder) Up() {
	fmt.Println("Seeder up started.")
	s.UserSeeder.Up()
	fmt.Println("Seeder up finished.")
}

func (s *AllSeeder) Down() {
	fmt.Println("Seeder down started.")
	s.UserSeeder.Down()
	fmt.Println("Seeder down finished.")
}
