package service

import (
	"YCSale/database"
	"YCSale/model"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func addWishService(m *martini.ClassicMartini) {
	m.Group("/wishes", func(router martini.Router) {
		router.Get("", func(r render.Render, db *mgo.Database) {
			r.HTML(200, "list", map[string]interface{}{
				"Title":  "Wish",
				"Wishes": model.WishModel.GetAll(db),
			})
		})
		router.Post("", binding.Form(model.Wish{}), func(w model.Wish, r render.Render, db *mgo.Database) {
			model.WishModel.Save(&w, db)
			r.HTML(200, "list", map[string]interface{}{
				"Title":  "Wish",
				"Wishes": model.WishModel.GetAll(db),
			})
		})
	}, database.GetMartini())
}
