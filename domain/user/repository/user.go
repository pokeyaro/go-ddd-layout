package repository

import (
	"server/domain/user/entity"
)

type UserRepository interface {
	UserRepositoryBase
	UserRepositoryQuery
}

type UserRepositoryBase interface {
	SaveUser(user *entity.User) (*entity.User, error)
	DeleteUser(id int64) error
	UpdateUser(user *entity.User) (*entity.User, error)
	UpdatesUser(user *entity.User) (*entity.User, error)
	GetUser(id int64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
}

type UserRepositoryQuery interface {
	FindByName(name string) (*entity.User, error)
	FindByEmpID(empId int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
	FindByUsernameAndPassword(username, password string) (*entity.User, error)
	UpdatePassword(id int64, password ...string) (string, error)
}
