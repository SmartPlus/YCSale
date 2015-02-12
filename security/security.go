package security

import (
	"YCSale/session"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
)

type SessionUser struct {
	Id    string
	Name  string
	Email string
	Role  int
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"Password"`
}

const (
	CURRENT_USER = "/current-user"
)

func SendSessionUser(r render.Render, s session.Session) {
	r.JSON(http.StatusOK, &SessionUser{
		Id:    session.String(s, "Id"),
		Name:  session.String(s, "Name"),
		Email: session.String(s, "Email"),
		Role:  session.Int(s, "Role"),
	})
}

func Init(m *martini.ClassicMartini) {
	m.Get(CURRENT_USER, SendSessionUser)
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
}
