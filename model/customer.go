package model

import (
	"gopkg.in/gorp.v1"
)

type Customer struct {
	Id         int    `json:"id"`
	LastName   string `json:"lastname"`
	FirstName  string `json:"firstname"`
	MiddleName string `json:"middlename"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Facebook   string `json:"facebook"`
	School     string `json:"school"`
	Company    string `json:"company"`
	Created_At string `json:"created_at"`
}

func (c *Customer) dbmapCustomer(dbmap *gorp.DbMap) {
	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	customerTab := dbmap.AddTableWithName(Customer{}, "customer")
	customerTab.ColMap("Created_At").SetTransient(true)
	customerTab.SetKeys(true, "Id")
}
