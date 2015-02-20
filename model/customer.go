package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Customer struct {
	Id       bson.ObjectId   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Phone    string          `json:"phone"`
	Facebook string          `json:"facebook"`
	School   string          `json:"school"`
	Company  string          `json:"company"`
	Courses  []bson.ObjectId `bson:"courses" json:"courses"`
	Notes    string          `json:"notes"`
}

type customerModel struct {
	Name string
}

var CustomerModel customerModel

func (m *customerModel) GetAll(db *mgo.Database) (customerList []Customer) {
	db.C(m.Name).Find(nil).All(&customerList)
	return
}

func (m *customerModel) Save(c *Customer, db *mgo.Database) {
	e := db.C(m.Name).Insert(c)
	if e != nil {
		println(e.Error())
	}
}

func (m *customerModel) RemoveById(id string, db *mgo.Database) error {
	return db.C(m.Name).RemoveId(bson.ObjectIdHex(id))
}

func (m *customerModel) FindById(id string, db *mgo.Database) (c *Customer) {
	c = &Customer{}
	db.C(m.Name).FindId(bson.ObjectIdHex(id)).One(c)
	return
}

func (m *customerModel) FindByEmail(email string, db *mgo.Database) (c *Customer, err error) {
	c = &Customer{}
	err = db.C(m.Name).Find(bson.M{"email": email}).One(c)
	return
}

func (m *customerModel) UpdateById(id string, c *Customer, db *mgo.Database) {
	c.Id = bson.ObjectIdHex(id) // safe user
	db.C(m.Name).UpdateId(c.Id, c)
	return
}

func (m *customerModel) Register(customerId string, courseId string, db *mgo.Database) {

	err := db.C(m.Name).Update(
		bson.M{
			"_id": bson.ObjectIdHex(customerId),
		},
		bson.M{
			"$addToSet": bson.M{
				"courses": bson.ObjectIdHex(courseId),
			},
		})

	if err != nil {
		println(err.Error())
	}
	return
}

func (m *customerModel) Unregister(customerId string, courseId string, db *mgo.Database) {

	err := db.C(m.Name).Update(
		bson.M{
			"_id": bson.ObjectIdHex(customerId),
		},
		bson.M{
			"$pull": bson.M{
				"courses": bson.ObjectIdHex(courseId),
			},
		})

	if err != nil {
		println(err.Error())
	}
	return
}

func (m *customerModel) Courses(customerId string, db *mgo.Database) (courseList []Course) {
	customer := m.FindById(customerId, db)
	return CourseModel.FindByIds(customer.Courses, db)
}
