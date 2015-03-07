package main

import (
	"YCSale/service"
	// "YCSale/security"
	"YCSale/session"
	"YCSale/web"
	// 	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	// "net/http"
)

const (
	COOKIE_SECRET      = "cki"
	COOKIE_KEY         = "my_session"
	TEMPLATE_DIRECTORY = "templates"
	TEMPLATE_LAYOUT    = "layout"
	ADMIN_PAGE         = "admin"
)

// func Authenticate() martini.Handler {
// 	return func(res http.ResponseWriter, req *http.Request, context martini.Context) {
// 		decoder := json.NewDecoder(req.Body)
// 		err := decoder.Decode(&security.SessionUser{})
// 		if err != nil {
// 			http.Error(res, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}
// 		context.Next()
// 	}
// }

func main() {
	m := martini.Classic()
	/* global middleware */
	m.Use(martini.Recovery())
	m.Use(martini.Static("www"))
	session.Init(COOKIE_SECRET, COOKIE_KEY, m)
	m.Use(render.Renderer(render.Options{
		Directory:  TEMPLATE_DIRECTORY,
		Layout:     TEMPLATE_LAYOUT,
		Extensions: []string{".html"},
	}))

	m.Use(service.DbMartiniHandler())

	web.Init(m)

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

	m.RunOnAddr(":8080")
}
