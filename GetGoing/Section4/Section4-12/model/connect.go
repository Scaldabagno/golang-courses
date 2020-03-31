package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

// Connect makes you connect to a db
func Connect() *sql.DB {
	log.Println("-> Connect()")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_example")
	err = db.Ping()
	if err != nil {
		log.Fatal("Error on Connect(): " + err.Error())
	}
	fmt.Println("Connected to db")
	con = db
	log.Println("<- Connect()")
	return db
}
