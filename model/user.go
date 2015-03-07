package model

type User struct {
	Id         int    `json:"id"`
	Role       string `json:"role"` /* Admin; CTV Basic; CTV Consultant; CSV Leader; CTV staff */
	LastName   string `json:"lastname"`
	FirstName  string `json:"firstname"`
	MiddleName string `json:"middlename"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Created_At string `json:"created_at"`
}
