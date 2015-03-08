package service

import (
	"YCSale/model"
	"github.com/go-martini/martini"
	"gopkg.in/gorp.v1"
	"strconv"
)

type UserService struct {
	dbmap *gorp.DbMap
}

func NewUserService(dbmap *gorp.DbMap) *UserService {
	return &UserService{dbmap: dbmap}
}

func (us *UserService) GetAll() (users []model.User) {
	_, err := us.dbmap.Select(&users, "SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	return
}

func (us *UserService) Insert(u *model.User) {
	err := us.dbmap.Insert(u)
	if err != nil {
		panic(err)
	}
}

func (us *UserService) Delete(s string) {
	id, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	_, err = us.dbmap.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
}

func (us *UserService) Update(u *model.User) {
	_, err := us.dbmap.Update(u)
	if err != nil {
		panic(err)
	}
}

func (us *UserService) Get(s string) (u *model.User) {
	u = &model.User{}
	id, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	err = us.dbmap.SelectOne(u, "SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	return
}

func (us *UserService) FindByEmail(email string) (u *model.User, err error) {
	u = &model.User{}
	err = us.dbmap.SelectOne(u, "SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return nil, err
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
