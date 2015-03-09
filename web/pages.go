package web

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

const (
	ADMIN_PAGE = "admin"
)

func addPages(m *martini.ClassicMartini) {
	m.Get("/admin", func(r render.Render) {
		r.HTML(200, ADMIN_PAGE, map[string]string{
			"Title": "Admin",
		})
	})
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", map[string]string{
			"Title": "YC Sale",
		})
	})
}
