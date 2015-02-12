package session

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

type Session sessions.Session

func String(session sessions.Session, key string) string {
	v := session.Get(key)
	if v == nil {
		return ""
	}
	return v.(string)
}

func Int(session sessions.Session, key string) int {
	v := session.Get(key)
	if v == nil {
		return -1
	}
	return v.(int)
}

func Init(secret, key string, m *martini.ClassicMartini) {
	store := sessions.NewCookieStore([]byte(secret))
	m.Use(sessions.Sessions(key, store))
}
