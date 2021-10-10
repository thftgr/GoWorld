package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
var Sqlite3 *sql.DB

func init()  {
	db, err := sql.Open("sqlite3", "./shortUrl.db")
	if err != nil {
		panic(err)
	}


	createTableQuery := `
		CREATE TABLE IF NOT EXISTS URLS (
			SHORT VARCHAR(10) NOT NULL PRIMARY KEY,
			URL TEXT NOT NULL,
		    SYSTEM_CREATE_TIMESTAMP TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`
	_, e := db.Exec( createTableQuery)
	if e != nil {
		panic(err)
	}

	Sqlite3 = db
}