package converter

import (
	"server/domain/user/entity"
	"server/infrastructure/persistence/user/po"
)

func UserEntityToPO(user *entity.User) *po.User {
	return &po.User{
		UUID:         user.UUID,
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

func UserPOToEntity(user *po.User) *entity.User {
	roles := make([]entity.Role, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = entity.Role{
			ID:       int(role.ID),
			RoleName: role.RoleName,
			RoleDesc: role.RoleDesc,
		}
	}

	return &entity.User{
		ID:           int(user.ID),
		UUID:         user.UUID,
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
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Roles:        roles,
	}
}
