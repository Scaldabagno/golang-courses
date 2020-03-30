package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

// Connect makes you connect to a db
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@(tcp:localhost:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to db")
	con = db
	return db
}
