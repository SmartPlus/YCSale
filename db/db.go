package db

import (
	"fmt"
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
	"net/http"
)

var (
	dbName   string
	host     string
	port     int
	username string
	password string
	url      string
	session  *mgo.Session
)

// DB Returns a martini.Handler
func Init(inUsername, inPassword, inHost string, inPort int, inDbName string) (err error) {
	dbName = inDbName
	username = inUsername
	password = inPassword
	host = inHost
	port = inPort

	url = fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", username, password, host, port, dbName)
	session, err = mgo.Dial(url)
	return
}

func Get() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		s := session.Clone()
		c.Map(s.DB(dbName))
		defer s.Close()
		c.Next()
	}
}
