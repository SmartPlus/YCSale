package first_run

import (
	"YCSale/model"
)

func createAdmin() *model.User {
	return &model.User{
		IUserContact: model.IUserContact{
			Role:      "admin",
			FirstName: "admin",
			LastName:  "admin",
			Email:     "admin@abc.com",
		},
		Password: "123",
		Phone:    "091",
	}
}
