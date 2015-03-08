package main

import (
	"YCSale/model"
	"YCSale/service"
	"fmt"
	"gopkg.in/gorp.v1"
	"log"
)

func main() {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := service.OpenDB(service.DB_NAME)
	checkErr(err, "Failed to open db")
	dbmap := service.GetDbMap(db)
	GetAllUser(dbmap)
	FindByEmail(dbmap, "admin@abc.com")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetAllUser(dbmap *gorp.DbMap) {
	u := service.NewUserService(dbmap)
	fmt.Printf("%v\n", u.GetAll())
}

func FindByEmail(dbmap *gorp.DbMap, email string) {
	var users []model.User
	info, err := dbmap.Select(&users, "SELECT * FROM user WHERE email=? LIMIT 1", email)
	fmt.Printf("%v\nInfo %v\nError %v\n", users, info, err)
}
