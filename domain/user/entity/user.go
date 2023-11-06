package entity

import (
	"time"

	"github.com/google/uuid"
)

// User Entity
type User struct {
	ID           int
	UUID         uuid.UUID
	Name         string
	NickName     string
	EmployeeID   int
	EmployeeType string
	Company      string
	WorkCountry  string
	WorkCity     string
	Department   string
	Email        string
	Avatar       string
	Username     string
	Password     string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Roles        Roles
}

// User Optional
type UserOpt func(*User)
type UserOpts []UserOpt

func (opts UserOpts) Apply(u *User) {
	for _, opt := range opts {
		opt(u)
	}
}

func WithID(id int) UserOpt {
	return func(u *User) {
		u.ID = id
	}
}

func WithName(name string) UserOpt {
	return func(u *User) {
		u.Name = name
	}
}

func WithEmployeeID(employeeID int) UserOpt {
	return func(u *User) {
		u.EmployeeID = employeeID
	}
}

func WithEmail(email string) UserOpt {
	return func(u *User) {
		u.Email = email
	}
}

func WithUsername(username string) UserOpt {
	return func(u *User) {
		u.Username = username
	}
}

func WithPassword(password string) UserOpt {
	return func(u *User) {
		u.Password = password
	}
}
