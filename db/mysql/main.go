package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/tests")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	query := "create table foo (id bigint auto_increment primary key, bar varchar(255))"
	if _, err := db.Exec(query); err != nil {
		panic(err)
	}

	query = "insert into foo (bar) values (?)"
	if _, err := db.Exec(query, "abcdef"); err != nil {
		panic(err)
	}
}