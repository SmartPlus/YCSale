package service

import (
	"YCSale/database"
	"YCSale/model"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func addCourseService(m *martini.ClassicMartini) {
	m.Group("/courses", func(router martini.Router) {
		router.Get("", func(r render.Render, db *mgo.Database) {
			r.JSON(200, model.CourseModel.GetAll(db))
		})
	}, database.GetMartini())

	m.Group("/course", func(router martini.Router) {
		router.Post("", binding.Json(model.Course{}), func(c model.Course, r render.Render, db *mgo.Database) {
			model.CourseModel.Save(&c, db)
		})
		router.Delete("/:_id", func(params martini.Params, r render.Render, db *mgo.Database) {
			model.CourseModel.RemoveById(params["_id"], db)
		})
		router.Get("/:_id", func(params martini.Params, r render.Render, db *mgo.Database) {
			r.JSON(200, model.CourseModel.FindById(params["_id"], db))
		})
		router.Put("/:_id", binding.Json(model.Course{}), func(params martini.Params, c model.Course, r render.Render, db *mgo.Database) {
			model.CourseModel.UpdateById(params["_id"], &c, db)
		})
	}, database.GetMartini())
}
