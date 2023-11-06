package user

type Userinfo struct {
	Name         string   `json:"name"`
	EmpID        int      `json:"emp_id"`
	Username     string   `json:"username"`
	Avatar       string   `json:"avatar"`
	Email        string   `json:"email"`
	Organization string   `json:"organization"`
	Company      string   `json:"company"`
	Region       string   `json:"region"`
	City         string   `json:"city"`
	Roles        []string `json:"roles"`
}

type User struct {
	ID           int    `json:"uid"`
	UUID         string `json:"uuid,omitempty"`
	Name         string `json:"name"`
	NickName     string `json:"nickname"`
	EmployeeID   int    `json:"employee_id,omitempty"`
	EmployeeType string `json:"employee_type,omitempty"`
	Company      string `json:"company,omitempty"`
	WorkCountry  string `json:"work_country,omitempty"`
	WorkCity     string `json:"work_city,omitempty"`
	Department   string `json:"department,omitempty"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	IsActive     bool   `json:"is_active"`
	CreatedAt    string `json:"ctime"`
	UpdatedAt    string `json:"utime"`
}

type LoginPasswd struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
