package service

import (
	"YCSale/model"
	"github.com/go-martini/martini"
)

func Init(m *martini.ClassicMartini, collections map[string]string) {
	model.Init(collections)
	addWishService(m)
}
