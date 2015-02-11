package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"gopkg.in/mgo.v2"
	"net/http"
)

type Wish struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}

type User struct {
	Id    string
	Name  string
	Email string
	Role  int
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"Password"`
}

type Learner struct {
	Id       string
	Name     string `json:"Name"`
	Phone    string `json:"Phone"`
	Facebook string
	Email    string
	School   string
	Company  string
	Courses  []string
	Notes    string
}

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

type Employee struct {
	Id   string
	Role int /* Admin; CTV Basic; CTV Consultant; CSV Leader; CTV staff */
	Name string
}

const (
	DATABASE_NAME   = "smartplus"
	COLLECTION_NAME = "wishes"
)

// DB Returns a martini.Handler
func DB() martini.Handler {
	session, err := mgo.Dial("mongodb://dong:123@ds031271.mongolab.com:31271/" + DATABASE_NAME)
	if err != nil {
		panic(err)
	}

	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		s := session.Clone()
		c.Map(s.DB(DATABASE_NAME))
		defer s.Close()
		c.Next()
	}
}

func Authenticate() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, context martini.Context) {
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&User{})
		if err != nil {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}
		context.Next()
	}
}

// GetAll returns all Wishes in the database
func GetAll(db *mgo.Database) map[string]interface{} {
	var wishlist []Wish
	db.C(COLLECTION_NAME).Find(nil).All(&wishlist)
	return map[string]interface{}{
		"Title":  "Wish",
		"Wishes": wishlist,
	}
}

func sessionString(session sessions.Session, key string) string {
	v := session.Get(key)
	if v == nil {
		return ""
	}
	return v.(string)
}

func sessionInt(session sessions.Session, key string) int {
	v := session.Get(key)
	if v == nil {
		return -1
	}
	return v.(int)
}

func sessionUser(r render.Render, session sessions.Session) {
	r.JSON(http.StatusOK, &User{
		Id:    sessionString(session, "Id"),
		Name:  sessionString(session, "Name"),
		Email: sessionString(session, "Email"),
		Role:  sessionInt(session, "Role"),
	})
}

func main() {
	m := martini.Classic()
	m.Use(martini.Static("www"))
	store := sessions.NewCookieStore([]byte("secret"))
	m.Use(sessions.Sessions("my_session", store))
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".html"},
	}))

	//START3 OMIT
	m.Use(DB())

	m.Get("/wishes", func(r render.Render, db *mgo.Database) {
		r.HTML(200, "list", GetAll(db))
	})

	m.Get("/admin", func(r render.Render, db *mgo.Database) {
		r.HTML(200, "admin", map[string]string{
			"Title": "Admin",
		})
	})

	m.Get("/current-user", sessionUser)
	m.Post("/login", binding.Json(LoginUser{}), func(user LoginUser, r render.Render) {
		if user.Email == "sp" && user.Password == "123" {
			r.JSON(200, map[string]interface{}{
				"user": map[string]string{
					"email":    "sp",
					"password": "123",
				},
			})
		} else {
			r.Error(http.StatusUnauthorized)
		}

	})
	m.Post("/wishes", binding.Form(Wish{}), func(wish Wish, r render.Render, db *mgo.Database) {
		db.C(COLLECTION_NAME).Insert(wish)
		r.HTML(200, "list", GetAll(db))
	})

	m.RunOnAddr(":8080")
}
