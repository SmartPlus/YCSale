package web

import (
	"YCSale/service"
	"YCSale/service/security"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
)

func addSecurityHandler(m *martini.ClassicMartini) {
	m.Group("/security", func(router martini.Router) {
		router.Get("/current-user", func(sec *security.SecurityService, r render.Render) {
			r.JSON(http.StatusOK, sec.CurrentUser())
		})
		router.Post("/login",
			binding.Json(security.LoginUser{}),
			service.UserMartiniHandler(),
			func(user security.LoginUser, us *service.UserService,
				sec *security.SecurityService, r render.Render,
				resp http.ResponseWriter) {
				err := sec.Login(&user, us)
				if err != nil {
					http.Error(resp, err.Error(), http.StatusForbidden)
				} else {
					r.JSON(http.StatusOK, sec.CurrentUser())
				}
			})
		router.Post("/logout", func(sec *security.SecurityService, resp http.ResponseWriter) {
			sec.Logout()
			resp.WriteHeader(http.StatusOK)
		})
	}, security.SecurityMartiniHandler())
}
