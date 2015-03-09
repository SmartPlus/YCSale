package web

import (
	"YCSale/service"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

const (
	WWW_DIRECTORY      = "www"
	TEMPLATE_DIRECTORY = "templates"
	TEMPLATE_LAYOUT    = "layout"
)

var www_dir = FilenameFromTheSameDir(WWW_DIRECTORY)
var template_dir = FilenameFromTheSameDir(TEMPLATE_DIRECTORY)

func Init() (m *martini.ClassicMartini) {
	m = martini.Classic()
	/* global middleware */
	m.Use(martini.Recovery())
	m.Use(martini.Static(www_dir))
	m.Use(render.Renderer(render.Options{
		Directory:  template_dir,
		Layout:     TEMPLATE_LAYOUT,
		Extensions: []string{".html"},
	}))
	service.Init(m)
	addPages(m)
	addHanders(m)
	return m
}

func addHanders(m *martini.ClassicMartini) {
	addUserHandlers(m)
	addSecurityHandler(m)
}
