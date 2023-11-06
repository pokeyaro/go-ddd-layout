package repository

import (
	"server/domain/user/entity"
)

type RoleRepository interface {
	SaveRole(role *entity.Role) error
	GetRole(id int64) (*entity.Role, error)
	GetRoles() (entity.Roles, error)
	SelectByRoleName(roleName string) (*entity.Role, error)
}
