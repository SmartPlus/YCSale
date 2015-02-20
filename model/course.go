package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Course struct {
	Id   bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Code string        `json:"code"`
	Fee  int           `json:"fee"`
	// Beginning Date
	// End Date
	// Payment Due
}

type courseModel struct {
	Name string
}

var CourseModel courseModel

func (m *courseModel) GetAll(db *mgo.Database) (courseList []Course) {
	db.C(m.Name).Find(nil).All(&courseList)
	return
}

func (m *courseModel) FindByIds(ids []bson.ObjectId, db *mgo.Database) (courseList []Course) {
	db.C(m.Name).Find(bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}).All(&courseList)
	return
}

func (m *courseModel) Save(c *Course, db *mgo.Database) {
	db.C(m.Name).Insert(c)
}

func (m *courseModel) RemoveById(id string, db *mgo.Database) error {
	return db.C(m.Name).RemoveId(bson.ObjectIdHex(id))
}

func (m *courseModel) FindById(id string, db *mgo.Database) (c *Course) {
	c = &Course{}
	db.C(m.Name).FindId(bson.ObjectIdHex(id)).One(c)
	return
}

func (m *courseModel) UpdateById(id string, c *Course, db *mgo.Database) {
	c.Id = bson.ObjectIdHex(id) // safe user
	db.C(m.Name).UpdateId(c.Id, c)
	return
}
