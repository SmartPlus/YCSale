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

func addCustomerService(m *martini.ClassicMartini) {
	m.Group("/customers", func(router martini.Router) {
		router.Get("", func(r render.Render, db *mgo.Database) {
			r.JSON(200, model.CustomerModel.GetAll(db))
		})
	}, security.RequireLoggedIn(), database.GetMartini())

	m.Group("/customer", func(router martini.Router) {
		router.Post("", binding.Json(model.Customer{}), func(c model.Customer, r render.Render, db *mgo.Database) {
			model.CustomerModel.Save(&c, db)
		})
		router.Delete("/:_id", func(params martini.Params, r render.Render, db *mgo.Database) {
			model.CustomerModel.RemoveById(params["_id"], db)
		})
		router.Get("/:_id", func(params martini.Params, r render.Render, db *mgo.Database) {
			r.JSON(200, model.CustomerModel.FindById(params["_id"], db))
		})
		router.Put("/:_id", binding.Json(model.Customer{}), func(params martini.Params, c model.Customer, r render.Render, db *mgo.Database) {
			model.CustomerModel.UpdateById(params["_id"], &c, db)
		})
		router.Get("/:_id/courses", func(params martini.Params, r render.Render, db *mgo.Database) {
			r.JSON(200, model.CustomerModel.Courses(params["_id"], db))
		})
		router.Post("/:customerId/register/:courseId", func(params martini.Params, r render.Render, db *mgo.Database) {
			model.CustomerModel.Register(params["customerId"], params["courseId"], db)
		})
		router.Post("/:customerId/unregister/:courseId", func(params martini.Params, r render.Render, db *mgo.Database) {
			model.CustomerModel.Unregister(params["customerId"], params["courseId"], db)
		})
	}, security.RequireLoggedIn(), database.GetMartini())
}
