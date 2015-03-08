package session

import (
	"YCSale/model"
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

const (
	COOKIE_SECRET = "s3kr3t"
	COOKIE_KEY    = "k3y"
	AUTH_KEY      = "authenticated"
	USER_KEY      = "user"
)

type User struct {
	model.IUserContact
	isLoggedIn bool
}

type SessionService struct {
	sessions.Session
	// cache user
}

func NewSessionService(s sessions.Session) *SessionService {
	return &SessionService{s}
}

func (s *SessionService) GetString(key string) string {
	v := s.Get(key)
	if v == nil {
		return ""
	}
	return v.(string)
}

func (s *SessionService) GetBool(key string) bool {
	v := s.Get(key)
	if v == nil {
		return false
	}
	return v.(bool)
}

func (s *SessionService) GetBytes(key string) []byte {
	v := s.Get(key)
	if v == nil {
		return nil
	}
	return v.([]byte)
}
func (s *SessionService) GetInt(key string) int {
	v := s.Get(key)
	if v == nil {
		return -1
	}
	return v.(int)
}

func (s *SessionService) IsAdmin() bool {
	return s.GetUser().IsAdmin()
}

func (s *SessionService) IsLoggedIn() bool {
	return s.GetBool(AUTH_KEY)
}

func (s *SessionService) Login() {
	s.Set(AUTH_KEY, true)
}

func (s *SessionService) Logout() {
	s.Delete(USER_KEY)
	s.Delete(AUTH_KEY)
}

func (s *SessionService) SetUser(u *User) {
	data, err := json.Marshal(u)
	if err != nil {
		return
	}
	s.Set(USER_KEY, data)
}

func (s *SessionService) GetUser() (u *User) {
	u = &User{}
	json.Unmarshal(s.GetBytes(USER_KEY), u)
	return
}

func Session() martini.Handler {
	store := sessions.NewCookieStore([]byte(COOKIE_SECRET))
	return sessions.Sessions(COOKIE_KEY, store)
}

func SessionMartiniHandler() martini.Handler {
	return func(s sessions.Session, c martini.Context) {
		c.Map(NewSessionService(s))
		c.Next()
	}
}
