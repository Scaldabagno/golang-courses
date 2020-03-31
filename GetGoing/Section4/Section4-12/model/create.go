package model

import "log"

// CreateTodo is the function for an insert query on the db
func CreateTodo(name, todo string) error {
	log.Println("-> CreateTodo")
	insertQ, err := con.Query("INSERT INTO todo VALUES(?, ?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		log.Println("Error on CreateTodo: " + err.Error())
		return err
	}
	log.Println("<- CreateTodo")
	return nil
}

// DeleteTodo is the function for a delete query on the db
func DeleteTodo(name string) error {
	log.Println("-> DeleteTodo")
	deleteQ, err := con.Query("DELETE FROM todo WHERE name=?", name)
	defer deleteQ.Close()
	if err != nil {
		log.Println("Error on DeleteTodo:" + err.Error())
		return err
	}
	log.Println("<- DeleteTodo")
	return nil
}
