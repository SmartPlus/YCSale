package security

import (
	"YCSale/database"
	"YCSale/model"
	"gopkg.in/mgo.v2"
)

const (
	ADMIN_EMAIL = "admin@abc.com"
)

func addAdminUser(db *mgo.Database) error {
	_, err := model.UserModel.FindByEmail(ADMIN_EMAIL, db)
	if err == mgo.ErrNotFound {
		err = model.UserModel.Save(&model.User{
			Name:     "Admin",
			Email:    ADMIN_EMAIL,
			Password: "123",
			Role:     "Admin",
		}, db)
	}
	return err
}

func InitDB() {
	database.WithDB(addAdminUser)
}
