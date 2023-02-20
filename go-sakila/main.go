package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root:tpc(localhost:3306)/sakila")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Print("Connected!")
}
