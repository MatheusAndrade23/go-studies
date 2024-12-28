package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		panic(err)
	}
	
	createTable := `
		CREATE TABLE foo (
			id integer not null primary key,
			name text
		);
	`

	res, err := db.Exec(createTable)

	if err != nil {
		panic(err)
	}

	fmt.Println(res.RowsAffected())

	insertSQL := `
		INSERT INTO foo (id, name) values (1, "Matheus")
	`

	res, err =  db.Exec(insertSQL)

	if err != nil {
		panic(err)
	}

	fmt.Println(res.RowsAffected())

	type user struct {
		ID int64
		Name string
	}

	querySQL := `
		SELECT * FROM foo WHERE id = ?;
	`
	var u user
	if err := db.QueryRow(querySQL, 1).Scan(&u.ID, &u.Name); err != nil {
		panic(err)
	}

	deleteSQL := `
		DELETE FROM foo WHERE id = ?
	`

	if _, err := db.Exec(deleteSQL); err != nil {
		panic(err)
	}
}