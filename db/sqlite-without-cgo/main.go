package main

import (
	"fmt"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

// _ "modernc.org/sqlite"

// func main() {
// 	db, err := sql.Open("sqlite", "./bar.foo")

// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := db.Ping(); err != nil {
// 		panic(err)
// 	}

// }

func main (){
	conn, err := sqlite.OpenConn("./baz.db", sqlite.OpenReadWrite | sqlite.OpenCreate)
	
	if err != nil {
  	panic(err) 
	}

	defer conn.Close()

	err = sqlitex.ExecuteTransient(conn, "SELECT 'hello, world';", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			fmt.Println(stmt.ColumnText(0))
			panic(err)
		},
	})

	if err != nil {
		panic(err)
	}
}