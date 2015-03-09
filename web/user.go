package web

import (
	"YCSale/model"
	"YCSale/service"
	"YCSale/service/security"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
)

func responseCheck(r render.Render, err error) {
	if err != nil {
		r.Data(http.StatusInternalServerError, []byte(err.Error()))
	} else {
		r.Status(http.StatusOK)
	}
}

func addUserHandlers(m *martini.ClassicMartini) {

	m.Group("/user", func(router martini.Router) {
		router.Get("/getall", func(r render.Render, us *service.UserService) {
			users, err := us.GetAll()
			if err != nil {
				responseCheck(r, err)
				return
			}
			r.JSON(http.StatusOK, users)
		})

		router.Post("/insert", binding.Json(model.User{}), func(u model.User, r render.Render, us *service.UserService) {
			responseCheck(r, us.Insert(&u))
		})

		router.Delete("/delete/:id", func(params martini.Params, r render.Render, us *service.UserService) {
			id, err := strconv.Atoi(params["id"])
			if err == nil {
				err = us.Delete(id)
			}
			responseCheck(r, err)
		})

		router.Get("/get/:id", func(params martini.Params, r render.Render, resp http.ResponseWriter, us *service.UserService) {
			id, err := strconv.Atoi(params["id"])
			if err == nil {
				if u, err := us.Get(id); err == nil {
					r.JSON(http.StatusOK, u)
					return
				}
			}
			responseCheck(r, err)
		})

		router.Put("/update/:id", binding.Json(model.User{}), func(u model.User, r render.Render, us *service.UserService) {
			err := us.Update(&u)
			responseCheck(r, err)
		})
	}, service.UserMartiniHandler(), security.SecurityMartiniHandler(), security.RequireAdmin())
}
