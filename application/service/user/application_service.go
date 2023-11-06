package user

import (
	"fmt"
	"strconv"

	"server/domain/user/entity"
	"server/domain/user/service"
	"server/infrastructure/common/jwt"
)

type Service interface {
	CreateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id int) error
	UpdateUser(user *entity.User, partialUpdate bool) (*entity.User, error)
	UpdatePassword(id int) (string, error)
	GetUserinfo(username string) (*entity.User, error)
	GetUserList() ([]entity.User, error)
	GetUserDetail(id int) (*entity.User, error)
	LocalLogin(username, password string) (string, error)
}

type ServiceImpl struct {
	ud service.UserDomain
}

func NewServiceImpl(srv service.UserDomain) Service {
	return &ServiceImpl{ud: srv}
}

func (srv *ServiceImpl) CreateUser(user *entity.User) (*entity.User, error) {
	return srv.ud.CreateUser(user)
}

func (srv *ServiceImpl) DeleteUser(id int) error {
	return srv.ud.DeleteUser(int64(id))
}

func (srv *ServiceImpl) UpdateUser(user *entity.User, partialUpdate bool) (*entity.User, error) {
	return srv.ud.UpdateUser(user, partialUpdate)
}

func (srv *ServiceImpl) UpdatePassword(id int) (string, error) {
	return srv.ud.UpdatePassword(int64(id))
}

func (srv *ServiceImpl) GetUserinfo(username string) (*entity.User, error) {
	return srv.ud.GetUserWithOpts(entity.WithUsername(username))
}

func (srv *ServiceImpl) GetUserList() ([]entity.User, error) {
	return srv.ud.GetUsers()
}

func (srv *ServiceImpl) GetUserDetail(id int) (*entity.User, error) {
	return srv.ud.GetUser(int64(id))
}

func (srv *ServiceImpl) LocalLogin(username, password string) (string, error) {
	user, err := srv.ud.GetUserWithOpts(
		entity.WithUsername(username),
		entity.WithPassword(password),
	)
	if err != nil {
		return "", fmt.Errorf("login failed: username '%s' and password do not match", username)
	}

	roleList := make([]string, 0, len(user.Roles))
	for _, role := range user.Roles {
		roleList = append(roleList, role.RoleName)
	}

	tokenData := jwt.TokenData{
		LoginUser:   user.Name,
		UserID:      user.ID,
		Roles:       roleList,
		EmpNO:       strconv.Itoa(user.EmployeeID),
		Avatar:      user.Avatar,
		AccessToken: user.UUID.String(),
	}

	token, err := jwt.CreateJwtToken(tokenData)
	if err != nil {
		return "", fmt.Errorf("failed to create JWT token: %w", err)
	}

	return token, nil
}
