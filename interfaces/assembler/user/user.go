package user

import (
	"time"

	"server/domain/user/entity"
	dto "server/interfaces/dto/user"
)

func DTOToEntity(user *dto.User) *entity.User {
	return &entity.User{
		Name:         user.Name,
		NickName:     user.NickName,
		EmployeeID:   user.EmployeeID,
		EmployeeType: user.EmployeeType,
		Company:      user.Company,
		WorkCountry:  user.WorkCountry,
		WorkCity:     user.WorkCity,
		Department:   user.Department,
		Email:        user.Email,
		Avatar:       user.Avatar,
		Username:     user.Username,
		Password:     user.Password,
		IsActive:     user.IsActive,
	}
}

func EntityToDTO(user *entity.User) *dto.User {
	return &dto.User{
		ID:           user.ID,
		UUID:         user.UUID.String(),
		Name:         user.Name,
		NickName:     user.NickName,
		EmployeeID:   user.EmployeeID,
		EmployeeType: user.EmployeeType,
		Company:      user.Company,
		WorkCountry:  user.WorkCountry,
		WorkCity:     user.WorkCity,
		Department:   user.Department,
		Email:        user.Email,
		Avatar:       user.Avatar,
		Username:     user.Username,
		Password:     user.Password,
		IsActive:     user.IsActive,
		CreatedAt:    user.CreatedAt.Format(time.DateTime),
		UpdatedAt:    user.UpdatedAt.Format(time.DateTime),
	}
}

func ListEntityToDTO(users []entity.User) []dto.User {
	var dtoUsers []dto.User
	for _, user := range users {
		dtoUser := dto.User{
			ID:           user.ID,
			Name:         user.Name,
			NickName:     user.NickName,
			EmployeeID:   user.EmployeeID,
			EmployeeType: user.EmployeeType,
			Company:      user.Company,
			WorkCountry:  user.WorkCountry,
			WorkCity:     user.WorkCity,
			Department:   user.Department,
			Email:        user.Email,
			Avatar:       user.Avatar,
			Username:     user.Username,
			Password:     user.Password,
			IsActive:     user.IsActive,
			CreatedAt:    user.CreatedAt.Format(time.DateTime),
			UpdatedAt:    user.UpdatedAt.Format(time.DateTime),
		}
		dtoUsers = append(dtoUsers, dtoUser)
	}
	return dtoUsers
}
