package service

import (
	"fmt"
	"strings"

	"server/domain/user/entity"
	"server/domain/user/repository"
)

type UserDomain interface {
	CreateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id int64) error
	UpdateUser(user *entity.User, partialUpdate bool) (*entity.User, error)
	UpdatePassword(id int64) (string, error)
	GetUser(id int64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserWithOpts(opts ...entity.UserOpt) (*entity.User, error)
}

type UserDomainImpl struct {
	ur repository.UserRepository
}

func NewUserDomainImpl(repo repository.UserRepository) UserDomain {
	return &UserDomainImpl{ur: repo}
}

func (ud *UserDomainImpl) CreateUser(user *entity.User) (*entity.User, error) {
	u, err := ud.ur.SaveUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, fmt.Errorf("username ‘%s’ already exists", user.Name)
		}
		return nil, err
	}
	return u, nil
}

func (ud *UserDomainImpl) DeleteUser(id int64) error {
	return ud.ur.DeleteUser(id)
}

func (ud *UserDomainImpl) UpdateUser(user *entity.User, partialUpdate bool) (*entity.User, error) {
	if partialUpdate {
		return ud.ur.UpdateUser(user)
	}
	return ud.ur.UpdatesUser(user)
}

func (ud *UserDomainImpl) UpdatePassword(id int64) (string, error) {
	return ud.ur.UpdatePassword(id)
}

func (ud *UserDomainImpl) GetUser(id int64) (*entity.User, error) {
	return ud.ur.GetUser(id)
}

func (ud *UserDomainImpl) GetUsers() ([]entity.User, error) {
	return ud.ur.GetUsers()
}

// GetUserWithOpts retrieves a user object based on the given options.
// The opts parameter is a series of UserOpt interfaces that customize the properties of the user object.
func (ud *UserDomainImpl) GetUserWithOpts(opts ...entity.UserOpt) (*entity.User, error) {
	user := new(entity.User)

	// Apply the options to the user object
	entity.UserOpts(opts).Apply(user)

	// Determine which fields are assigned values
	hasFieldName := user.Name != ""
	hasFieldUsername := user.Username != ""
	hasFieldPassword := user.Password != ""
	hasFieldEmployeeID := user.EmployeeID != 0
	hasFieldEmail := user.Email != ""

	// Perform corresponding actions based on different conditions
	switch {
	case hasFieldName:
		return ud.ur.FindByName(user.Name)
	case hasFieldUsername:
		if hasFieldPassword {
			return ud.ur.FindByUsernameAndPassword(user.Username, user.Password)
		}
		return ud.ur.FindByUsername(user.Username)
	case hasFieldEmployeeID:
		return ud.ur.FindByEmpID(user.EmployeeID)
	case hasFieldEmail:
		return ud.ur.FindByEmail(user.Email)
	default:
		return ud.ur.GetUser(int64(user.ID))
	}
}
