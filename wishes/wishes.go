package wishes

import (
	"YCSale/db"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func Init(m *martini.ClassicMartini, col string) {
	setName(col)
	m.Group("/wishes", func(router martini.Router) {
		router.Get("", func(r render.Render, db *mgo.Database) {
			r.HTML(200, "list", map[string]interface{}{
				"Title":  "Wish",
				"Wishes": GetAll(db),
			})
		})
		router.Post("", binding.Form(Wish{}), func(wish Wish, r render.Render, db *mgo.Database) {
			(&wish).Save(db)
			r.HTML(200, "list", map[string]interface{}{
				"Title":  "Wish",
				"Wishes": GetAll(db),
			})
		})
	}, db.Get())
}
