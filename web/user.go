package web

import (
	"YCSale/model"
	"YCSale/service"
	"YCSale/service/security"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func addUserHandlers(m *martini.ClassicMartini) {

	m.Group("/user", func(router martini.Router) {
		router.Get("/getall", func(r render.Render, us *service.UserService) {
			r.JSON(200, us.GetAll())
		})

		router.Post("/insert", binding.Json(model.User{}), func(u model.User, r render.Render, us *service.UserService) {
			us.Insert(&u)
		})

		router.Delete("/delete/:_id", func(params martini.Params, r render.Render, us *service.UserService) {
			println("Id: " + params["_id"])
			us.Delete(params["_id"])
		})

		router.Get("/get/:_id", func(params martini.Params, r render.Render, us *service.UserService) {
			r.JSON(200, us.Get(params["_id"]))
		})

		router.Put("/update/:_id", binding.Json(model.User{}), func(params martini.Params, u model.User, r render.Render, us *service.UserService) {
			us.Update(&u)
		})
	}, service.UserMartiniHandler(), security.SecurityMartiniHandler(), security.RequireAdmin())
}
