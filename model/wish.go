package model

import (
	"gopkg.in/gorp.v1"
)

type Wish struct {
	Id          string
	Name        string `form:"name"`
	Description string `form:"description"`
}

func (w *Wish) dbmapWish(dbmap *gorp.DbMap) {
	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	wishTab := dbmap.AddTableWithName(Wish{}, "wish")
	wishTab.ColMap("Name").SetMaxSize(128)
	wishTab.SetKeys(true, "Id")
}
