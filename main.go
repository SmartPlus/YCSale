package main

import (
	"YCSale/database"
	"YCSale/security"
	"YCSale/service"
	"YCSale/session"
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type Course struct {
	Id   string
	Name string
	Cost int
}

type Payment struct {
	Id      string
	Course  string
	Learner string
	Ammount int
}

const (
	MONGO_USERNAME            = "debug" //"dong"
	MONGO_PASSWORD            = "123"
	MONGO_HOST                = "localhost" //"ds031271.mongolab.com"
	MONGO_PORT                = 27017       //31271
	MONG_DB_NAME              = "smartplus"
	WISHES_COLLECTION_NAME    = "wishes"
	USERS_COLLECTION_NAME     = "users"
	CUSTOMERS_COLLECTION_NAME = "customers"
	COURSES_COLLECTION_NAME   = "courses"
	COOKIE_SECRET             = "cki"
	COOKIE_KEY                = "my_session"
	TEMPLATE_DIRECTORY        = "templates"
	TEMPLATE_LAYOUT           = "layout"
	ADMIN_PAGE                = "admin"
)

func Authenticate() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, context martini.Context) {
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&security.SessionUser{})
		if err != nil {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}
		context.Next()
	}
}

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

	if err := database.Init(MONGO_USERNAME, MONGO_PASSWORD, MONGO_HOST, MONGO_PORT, MONG_DB_NAME); err != nil {
		panic(err)
	}

	service.Init(m, map[string]string{
		"wish":     WISHES_COLLECTION_NAME,
		"user":     USERS_COLLECTION_NAME,
		"customer": CUSTOMERS_COLLECTION_NAME,
		"course":   COURSES_COLLECTION_NAME,
	})

	// security use service
	security.Init(m)
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
