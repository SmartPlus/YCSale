package security

import (
	"YCSale/database"
	"YCSale/model"
	"YCSale/session"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"net/http"
)

type SessionUser struct {
	Id    string `json:"_id,omitempty"`
	Role  string `json:"role"` /* Admin; CTV Basic; CTV Consultant; CSV Leader; CTV staff */
	Name  string `json:"name"`
	Email string `json:"email"`
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
		Role:  session.String(s, "Role"),
	})
}

func Init(m *martini.ClassicMartini) {
	m.Get(CURRENT_USER, SendSessionUser)
	m.Post("/login", binding.Json(LoginUser{}), database.GetMartini(), func(user LoginUser, r render.Render, db *mgo.Database) {
		u := model.UserModel.FindByEmail(user.Email, db)
		if u.Email != user.Email {
			r.Data(http.StatusUnauthorized, []byte("Email is not existed"))
			return
		}

		if u.Password != user.Password {
			r.Data(http.StatusUnauthorized, []byte("Wrong password"))
			return
		}

		s := SessionUser{
			Id:    u.Id.Hex(),
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		}

		r.JSON(200, map[string]interface{}{
			"user": s,
		})
	})
}
