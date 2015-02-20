package service

import (
	"YCSale/database"
	"YCSale/model"
	"YCSale/security"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func addUserService(m *martini.ClassicMartini) {
	m.Group("/users", func(router martini.Router) {
		router.Get("", func(r render.Render, db *mgo.Database) {
			r.JSON(200, model.UserModel.GetAll(db))
		})
	}, security.RequireAdmin(), database.GetMartini())

	m.Group("/user", func(router martini.Router) {
		router.Post("", binding.Json(model.User{}), func(u model.User, r render.Render, db *mgo.Database) {
			model.UserModel.Save(&u, db)
		})
		router.Delete("/:_id", func(params martini.Params, r render.Render, db *mgo.Database) {
			println("Id: " + params["_id"])
			model.UserModel.RemoveById(params["_id"], db)
		})
		router.Get("/:_id", func(params martini.Params, r render.Render, db *mgo.Database) {
			r.JSON(200, model.UserModel.FindById(params["_id"], db))
		})
		router.Put("/:_id", binding.Json(model.User{}), func(params martini.Params, u model.User, r render.Render, db *mgo.Database) {
			model.UserModel.UpdateById(params["_id"], &u, db)
		})
	}, security.RequireAdmin(), database.GetMartini())
}
