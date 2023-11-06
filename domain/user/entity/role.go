package entity

// Role Entity
type Role struct {
	ID       int    `json:"id,omitempty"`
	RoleName string `json:"role_name,omitempty"`
	RoleDesc string `json:"role_desc,omitempty"`
}

type Roles []Role
