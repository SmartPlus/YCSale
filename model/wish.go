package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Wish struct {
	Id          bson.ObjectId
	Name        string `form:"name"`
	Description string `form:"description"`
}

type wishModel struct {
	Name string
}

var WishModel wishModel

/*
func SearchWish(q interface{}, skip int, limit int) (searchResults []Wish, searchErr string) {
	searchErr = ""
	searchResults = []Wish{}
	query := func(c *mgo.Collection) error {
		fn := c.Find(q).Skip(skip).Limit(limit).All(&searchResults)
		if limit < 0 {
			fn = c.Find(q).Skip(skip).All(&searchResults)
		}
		return fn
	}

	search := func() error {
		return WithCollection("wishes", query)
	}
	err := search()
	if err != nil {
		searchErr = "Database Error"
	}
	return
}

func GetWishByName(name string, skip int, limit int) (searchResults []Wish, searchErr string) {
	searchResults, searchErr = SearchWish(bson.M{"name": bson.RegEx{"^" + name, "i"}}, skip, limit)
	return
}

*/
func (m *wishModel) GetAll(db *mgo.Database) []Wish {
	var wishlist []Wish
	db.C(m.Name).Find(nil).All(&wishlist)
	return wishlist
}

func (m *wishModel) Save(w *Wish, db *mgo.Database) {
	db.C(m.Name).Insert(w)
}
