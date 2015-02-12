package wishes

import (
	"gopkg.in/mgo.v2"
)

type Wish struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}

var (
	collectionName string
)

// GetAll returns all Wishes in the database
func GetAll(db *mgo.Database) []Wish {
	var wishlist []Wish
	db.C(collectionName).Find(nil).All(&wishlist)
	return wishlist
}

func setName(name string) {
	collectionName = name
}

func (w *Wish) Save(db *mgo.Database) {
	db.C(collectionName).Insert(w)
}
