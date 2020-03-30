package model

import (
	"../views"
	"log"
)

func ReadAll() ([]views.PostRequest, error) {
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
	return todos, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	log.Println("-> ReadByName()")
	rows, err := con.Query("SELECT * FROM todo WHERE name=?", name)
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
