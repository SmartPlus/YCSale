package service

import (
	"YCSale/model"
	"database/sql"
	"errors"
	"github.com/go-martini/martini"
	"gopkg.in/gorp.v1"
)

type UserService struct {
	dbmap *gorp.DbMap
}

func NewUserService(dbmap *gorp.DbMap) *UserService {
	return &UserService{dbmap: dbmap}
}

func (us *UserService) GetAll() (users []model.User, err error) {
	_, err = us.dbmap.Select(&users, "SELECT * FROM user")
	if err != nil {
		return nil, LogError(err)
	}
	return users, nil
}

func (us *UserService) Insert(u *model.User) (err error) {
	u0, err := us.FindByEmail(u.Email)
	if err != nil {
		return
	}

	if u0 != nil {
		return errors.New("Email existed")
	}

	err = us.dbmap.Insert(u)
	return LogError(err)
}

func (us *UserService) Delete(id int) error {
	_, err := us.dbmap.Exec("DELETE FROM user WHERE id = ?", id)
	return LogError(err)
}

func (us *UserService) Update(u *model.User) error {
	_, err := us.dbmap.Update(u)
	if err != nil {
		return LogError(err)
	}
	return nil
}

func (us *UserService) Get(id int) (u *model.User, err error) {
	u = &model.User{}

	err = us.dbmap.SelectOne(u, "SELECT * FROM user WHERE id = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, LogError(err)
		}
	}
	return
}

func (us *UserService) FindByEmail(email string) (u *model.User, err error) {
	u = &model.User{}
	err = us.dbmap.SelectOne(u, "SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, LogError(err)
		}
	}
	return
}

func (us *UserService) Map() {
	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	userTab := us.dbmap.AddTableWithName(model.User{}, "user")
	userTab.ColMap("Created_At").SetTransient(true)
	userTab.SetKeys(true, "Id")
}

func UserMartiniHandler() martini.Handler {
	return func(dbmap *gorp.DbMap, c martini.Context) {
		c.Map(NewUserService(dbmap))
		c.Next()
	}
}
