package mylibaries

import (
	"database/sql"
)

func mySQL(dbHost string, dbUser string, dbPass string, dbName string) (db *sql.DB) {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@"+dbHost+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
