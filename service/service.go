package service

import (
	"YCSale/service/session"
	"github.com/go-martini/martini"
)

func Init(m *martini.ClassicMartini) {
	m.Use(session.Session())
	m.Use(DbMartiniHandler())
}
