package main

import (
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"log"
)

func init() {
	db, err := sql.Open("sqlite3", "./zoltan.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`create table if not exists users (id integer not null primary key autoincrement,
		username text, hash blob);`)
	check(err)
}

func Login(username string, password string) (bool, string) {
	return false, "Authentication failure"
}

func Register(username string, password string) (bool, string) {
	return false, ""
}
