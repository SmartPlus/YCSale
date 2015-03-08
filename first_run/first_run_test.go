package first_run

import (
	"YCSale/model"
	"YCSale/service"
	"fmt"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	InitDB()
}

func TestInsertAdmin(t *testing.T) {
	db, err := service.OpenDB(service.DB_NAME)
	checkErr(err, "Failed to open db")
	dbmap := service.GetDbMap(db)
	expected := createAdmin()
	admin := &model.User{}
	err = dbmap.SelectOne(admin, "SELECT * FROM user WHERE email=?", "admin@abc.com")
	checkErr(err, "read admin user")
	fmt.Printf("%+v\n %+v\n", admin, expected)
	admin.Id = expected.Id
	admin.Created_At = expected.Created_At
	if !reflect.DeepEqual(admin, expected) {
		t.Error("Insert Admin Fail!")
	}
}
