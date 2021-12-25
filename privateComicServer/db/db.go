package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var Sqlite3 *sql.DB
var Token = map[string]struct{
	Id int
	Admin bool
}{}



func init() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		panic(err)
	}

	createTableQuery :=
		`
-- 유저 테이블
CREATE TABLE IF NOT EXISTS USER(
	ID INTEGER NOT NULL CONSTRAINT USER_PK PRIMARY KEY AUTOINCREMENT,
	LOGIN_ID VARCHAR(100) NOT NULL,
	PASSWORD VARCHAR(63) NOT NULL,
	EMAIL VARCHAR(1024) NOT NULL,
	GROUP_ID INTEGER NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS USER_ID_UINDEX
	ON USER (ID);

CREATE UNIQUE INDEX IF NOT EXISTS USER_LOGIN_ID_UINDEX
	ON USER (LOGIN_ID);

`
	_, e := db.Exec(createTableQuery)
	if e != nil {
		panic(e)
	}

	Sqlite3 = db
}
