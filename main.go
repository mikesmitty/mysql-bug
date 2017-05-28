package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Get our connection details
	f, err := ioutil.ReadFile("dsn.txt")
	if err != nil {
		panic(err)
	}
	dsn := string(bytes.TrimSpace(f))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Problems exhibited when connection lifetimes are set
	db.SetConnMaxLifetime(2 * time.Second)

	var id int
	for i := 0; i < 3; i++ {
		rows, err := db.Query("SELECT * FROM bugtest WHERE id < ?", i)
		for rows.Next() {
			err = rows.Scan(&id)
			if err != nil {
				panic(err)
			}
		}
		rows.Close()

		fmt.Printf("Run #%d\n", i+1)
		time.Sleep(1 * time.Second)
	}
}
