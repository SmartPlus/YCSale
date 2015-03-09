package model

type User struct {
	IContact
	IUser
	Password   string `json:"-"`
	Created_At string `json:"-"`
}

type IUser struct {
	Role string `json:"role"` /* Admin; CTV Basic; CTV Consultant; CSV Leader; CTV staff */
}

type IContact struct {
	Id         int    `json:"id,omitempty"`
	LastName   string `json:"lastname"`
	FirstName  string `json:"firstname"`
	MiddleName string `json:"middlename"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

func (u *IUser) IsAdmin() bool {
	return u.Role == "admin"
}
