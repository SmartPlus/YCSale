package service

import (
	"errors"
	"log"
	"os"
)

var errLog = log.New(os.Stderr, "[Service] ", log.Ldate|log.Ltime|log.Lshortfile)
var (
	database_error = errors.New("Database error")
)

/*
	Log the error
	Return a database error for public user
*/

func LogError(err error) error {
	if err != nil {
		errLog.Print(err.Error())
		return database_error
	}
	return nil
}
