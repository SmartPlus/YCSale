package service

import (
	"database/sql"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"io/ioutil"
	"log"
	"strings"
)

const (
	CREATE_DATABASE_SQL     = "./create_database.sql"
	SQL_STATEMENT_SEPERATOR = ";"
	DB_NAME                 = "smartplus"
)

func OpenDB(dbname string) (db *sql.DB, err error) {
	return sql.Open("mysql", "admin:123456@tcp(localhost:3306)/"+dbname)
}

func GetDbMap(db *sql.DB) *gorp.DbMap {
	// construct a gorp DbMapt
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.TraceOn("[MySQL]", errLog)

	NewUserService(dbmap).Map()
	// (&Customer{}).dbmapCustomer(dbmap)
	// (&Course{}).dbmapCourse(dbmap)
	// (&Student{}).dbmapStudent(dbmap)
	// (&Wish{}).dbmapWish(dbmap)
	dbmap.CreateTablesIfNotExists()
	return dbmap
}

func ExecFile(db *sql.DB, filename, seperator string) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	queries := strings.Split(string(contents), seperator)
	for _, query := range queries {
		if len(query) > 0 {
			_, err := db.Exec(query)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func DbMartiniHandler() martini.Handler {
	db, err := OpenDB(DB_NAME)
	checkErr(err, "Failed to open db")
	dbmap := GetDbMap(db)

	return func(c martini.Context) {
		c.Map(dbmap)
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
