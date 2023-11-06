package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id;            type:int;          not null; primaryKey; autoIncrement;    comment:'Primary key ID'"`
	RoleName string `gorm:"column:role_name;     type:varchar(36);  index:idx_role_name; unique; not null;  comment:'Role name'"`
	RoleDesc string `gorm:"column:role_desc;     type:text;                                                 comment:'Role description'"`
}

func (u *Role) TableName() string {
	return "t_rbac_role"
}
