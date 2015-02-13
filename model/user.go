package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Role     string        `json:"role"` /* Admin; CTV Basic; CTV Consultant; CSV Leader; CTV staff */
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
}

type userModel struct {
	Name string
}

var UserModel userModel

func (m *userModel) GetAll(db *mgo.Database) (userlist []User) {
	db.C(m.Name).Find(nil).All(&userlist)
	return
}

func (m *userModel) Save(u *User, db *mgo.Database) {
	db.C(m.Name).Insert(u)
}

func (m *userModel) RemoveById(id string, db *mgo.Database) {
	err := db.C(m.Name).RemoveId(bson.ObjectIdHex(id))
	println(err.Error())
}

func (m *userModel) FindById(id string, db *mgo.Database) (u User) {
	db.C(m.Name).FindId(bson.ObjectIdHex(id)).One(&u)
	return
}

func (m *userModel) UpdateById(id string, u *User, db *mgo.Database) {
	u.Id = bson.ObjectIdHex(id) // safe user
	db.C(m.Name).UpdateId(bson.ObjectIdHex(id), u)
	return
}
