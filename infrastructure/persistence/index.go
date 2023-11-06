package persistence

import (
	"server/domain/user/repository"
	"server/infrastructure/persistence/user"

	"gorm.io/gorm"
)

type Repositories struct {
	User repository.UserRepository
	// other repositories...
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewUserRepository(db),
	}
}
