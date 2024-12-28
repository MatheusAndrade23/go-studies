package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	urlExample := "postgres://pg:password@localhost:5432/tests"

	// db, err := pgx.Connect(context.Background(), urlExample)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer db.Close(context.Background())

	db, err := pgxpool.New(context.Background(), urlExample)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	query := "create table if not exists foo (id bigserial primary key, bar varchar(255))"
	if _, err := db.Exec(context.Background(), query); err != nil {
		panic(err)
	}

	query = "insert into foo (bar) values ($1)"
	if _, err := db.Exec(context.Background(), query, "abcdef"); err != nil {
		panic(err)
	}
}