package model

import (
	"gopkg.in/gorp.v1"
)

type Student struct {
	Id          int    `json:"id"`
	Customer_id int    `json:"customer_id"`
	Course_Id   int    `json:"course_id"`
	Paid_Amount int    `json:"paid_amount"`
	Created_At  string `json:"created_at"`
}

func (s *Student) dbmapStudent(dbmap *gorp.DbMap) {
	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	studentTab := dbmap.AddTableWithName(Student{}, "student")
	studentTab.ColMap("Created_At").SetTransient(true)
	studentTab.SetKeys(true, "Id")
}
