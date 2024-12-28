package main

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

func main() {
	filters := Filters{
		ID:       123,
		Name:     "abc",
		Email:    "abc",
		Username: "abc",
	}

	fmt.Println(build(filters))
}

type Filters struct {
	ID       int64
	Name     string
	Email    string
	Username string
}

func build(f Filters) (string, []any) {
	// builder := squirrel.Select("*").From("Users") -> MySQL placeholders "?"
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Select("*").From("Users") // Postgres placeholder "$"
	or := squirrel.Or{}

	if f.ID > 0 {
		or = append(or, squirrel.Eq{"id": f.ID})
	}

	if f.Name != "" {
		or = append(or, squirrel.Eq{"name": "%" + f.Name + "%"})
	}

	if f.Email != "" {
		or = append(or, squirrel.Eq{"email": f.Email})
	}

	if f.Username != "" {
		or = append(or, squirrel.Eq{"username": f.Username})
	}

	sql, args, err := builder.Where(or).ToSql()

	if err != nil {
		panic(err)
	}

	return sql, args
}