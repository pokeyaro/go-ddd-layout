package po

import (
	"server/domain/user/entity/valueobj"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID         uuid.UUID `gorm:"column:uuid;           type:varchar(36);   not null; index:idx_uuid; unique;     comment:'Unique identifier'"`
	Name         string    `gorm:"column:name;           type:varchar(50);   not null; index:idx_name;             comment:'English name'"`
	NickName     string    `gorm:"column:nickname;       type:varchar(50);   not null;                             comment:'Chinese name'"`
	EmployeeID   int       `gorm:"column:employee_id;    type:int;                                                 comment:'Employee ID'"`
	EmployeeType string    `gorm:"column:employee_type;  type:varchar(50);                                         comment:'Employee type'"`
	Company      string    `gorm:"column:company;        type:varchar(50);                                         comment:'Company name'"`
	WorkCountry  string    `gorm:"column:work_country;   type:varchar(50);                                         comment:'Work country'"`
	WorkCity     string    `gorm:"column:work_city;      type:varchar(50);                                         comment:'Work city'"`
	Department   string    `gorm:"column:department;     type:varchar(50);                                         comment:'Department'"`
	Email        string    `gorm:"column:email;          type:varchar(50);   not null; index:idx_email; unique;    comment:'Email address'"`
	Avatar       string    `gorm:"column:avatar;         type:text;                                                comment:'User avatar URL'"`
	Username     string    `gorm:"column:username;       type:varchar(50);   index:idx_username; unique;           comment:'System login username'"`
	Password     string    `gorm:"column:password;       type:varchar(255);  index:idx_password;                   comment:'Local system login password'"`
	IsActive     bool      `gorm:"column:is_active;      type:boolean;       index:idx_is_active; default:1;       comment:'Whether the user is active (0 for inactive, 1 for active)'"`
	Roles        []Role    `gorm:"many2many:rbac_user_roles;"`
}

func (u *User) TableName() string {
	return "t_rbac_user"
}

// BeforeCreate - Pre Insert Hook.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.IsActive = true

	u.UUID = uuid.New()

	var role Role
	tx.Where("role_name = ?", valueobj.DefaultRoleName).First(&role)
	u.Roles = []Role{role}

	return
}
