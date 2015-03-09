package security

import (
	"YCSale/service"
	"YCSale/service/session"
	"errors"
	"github.com/go-martini/martini"
	session_martini "github.com/martini-contrib/sessions"
	"net/http"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"Password"`
}

type SecurityService struct {
	session *session.SessionService
}

func NewSecurityService(ss *session.SessionService) *SecurityService {
	return &SecurityService{session: ss}
}

func (sec *SecurityService) CurrentUser() map[string]interface{} {
	if sec.session.IsLoggedIn() {
		return map[string]interface{}{
			"user": sec.session.GetUser(),
		}
	}
	return map[string]interface{}{}
}

func (sec *SecurityService) Login(loginUser *LoginUser, us *service.UserService) (err error) {
	user, err := us.FindByEmail(loginUser.Email)

	if err != nil {
		return errors.New("Unknown Email")
	}

	if user.Email == "" {
		return errors.New("Unknown Email")
	}

	if user.Password != loginUser.Password {
		return errors.New("Wrong Password")
	}

	sessionUser := session.NewSessionUser(user)
	sec.session.SetUser(sessionUser)
	sec.session.Login()
	return nil
}

func (sec *SecurityService) Logout() {
	sec.session.Logout()
}

func RequireAdmin() martini.Handler {
	return func(resp http.ResponseWriter, sec *SecurityService, c martini.Context) {
		if sec.session.IsAdmin() {
			c.Next()
		} else {
			resp.WriteHeader(http.StatusUnauthorized)
		}
	}
}

func RequireLoggedIn() martini.Handler {
	return func(resp http.ResponseWriter, sec *SecurityService, c martini.Context) {
		if sec.session.IsLoggedIn() {
			c.Next()
		} else {
			resp.WriteHeader(http.StatusUnauthorized)
		}
	}
}

func SecurityMartiniHandler() martini.Handler {
	return func(s session_martini.Session, c martini.Context) {
		c.Map(NewSecurityService(session.NewSessionService(s)))
		c.Next()
	}
}
