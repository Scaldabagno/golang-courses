package model

import (
	"log"

	"../views"
)

// ReadAll is the function for a select * query on the db
func ReadAll() ([]views.PostRequest, error) {
	log.Println("-> ReadAll()")
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	log.Println("<- ReadAll()")
	return todos, nil
}

// ReadByName is the function for a select * where query on the db
func ReadByName(name string) ([]views.PostRequest, error) {
	log.Println("-> ReadByName()")
	rows, err := con.Query("SELECT * FROM TODO WHERE name=?", name)
	if err != nil {
		log.Println("Error on ReadByName()")
		return nil, err
	}
	// var todos []views.PostRequest
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	log.Println("<- ReadByName()")
	return todos, nil
}
