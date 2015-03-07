package first_run

import (
	"YCSale/service"
	"log"
)

func InitDB() {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := service.OpenDB("")
	checkErr(err, "Failed to open db")
	dbmap := service.GetDbMap(db)
	service.ExecFile(db, service.CREATE_DATABASE_SQL, service.SQL_STATEMENT_SEPERATOR)
	checkErr(dbmap.Insert(createAdmin()), "Failed to insert admin user")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
