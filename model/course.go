package model

import (
	"gopkg.in/gorp.v1"
)

type Course struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Teacher     int    `json:"teacher"`
	Fee         int    `json:"fee"`
	Timetable   string `json:"timetable"`
	Start_Date  string `json:"start_date"`
	End_Date    string `json:"end_date"`
	Created_At  string `json:"created_at"`
}

func (c *Course) dbmapCourse(dbmap *gorp.DbMap) {
	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	courseTab := dbmap.AddTableWithName(Course{}, "course")
	courseTab.ColMap("Created_At").SetTransient(true)
	courseTab.SetKeys(true, "Id")
}
