package database

import (
	"fmt"
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
	"net/http"
)

type DB struct {
	username string
	password string
	host     string
	port     int
	database string
	url      string
	session  *mgo.Session
}

var context DB

// DB Returns a martini.Handler
func Init(username, password, host string, port int, database string) (err error) {
	var session *mgo.Session
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", username, password, host, port, database)
	session, err = mgo.Dial(url)
	if err != nil {
		return
	}

	context = DB{
		username: username,
		password: password,
		host:     host,
		port:     port,
		database: database,
		url:      url,
		session:  session,
	}
	return
}

func GetMartini() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		s := context.session.Clone()
		c.Map(s.DB(context.database))
		defer s.Close()
		c.Next()
	}
}

func WithCollection(collection string, s func(*mgo.Collection) error) error {
	session := context.session.Clone()
	defer session.Close()
	c := session.DB(context.database).C(collection)
	return s(c)
}
