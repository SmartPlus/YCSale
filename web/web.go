package web

import (
	"github.com/go-martini/martini"
)

func Init(m *martini.ClassicMartini) {
	addUserHandlers(m)
}
