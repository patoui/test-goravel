package seeders

import (
	"fmt"
	"goravel/app/models"
	"math/rand"

	"github.com/goravel/framework/facades"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	user := models.User{
		Name: "goravel",
		Email: fmt.Sprintf("goravel+%d@test.com", rand.Intn(100)),
	}
	return facades.Orm().Query().Create(&user)
}
