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

func SetSessionUser(u SessionUser, s session.Session) {
	s.Set("Id", u.Id)
	s.Set("Name", u.Name)
	s.Set("Email", u.Email)
	s.Set("Role", u.Role)
}

func Unauthorized(r render.Render, s string) {
	r.Data(402, []byte(s))
}

func Init(m *martini.ClassicMartini) {
	m.Get(CURRENT_USER, SendSessionUser)
	m.Post("/login", binding.Json(LoginUser{}), database.GetMartini(), func(user LoginUser, r render.Render, s session.Session, db *mgo.Database) {
		u, err := model.UserModel.FindByEmail(user.Email, db)
		if err != nil {
			if err == mgo.ErrNotFound {
				Unauthorized(r, "Email is not existed")
			} else {
				Unauthorized(r, "Database error")
			}
			return
		}

		if u.Password != user.Password {
			Unauthorized(r, "Wrong password")
			return
		}

		sUser := SessionUser{
			Id:    u.Id.Hex(),
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		}
		SetSessionUser(sUser, s)
		r.JSON(200, map[string]interface{}{
			"user": sUser,
		})
	})

	InitDB()
}
