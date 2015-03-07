package first_run

import (
	"YCSale/model"
)

func createAdmin() *model.User {
	return &model.User{
		Role:      "admin",
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin@smartplus.com",
		Password:  "123456",
		Phone:     "091",
	}
}
