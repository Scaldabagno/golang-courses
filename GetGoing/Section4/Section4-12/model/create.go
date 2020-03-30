package model

// CreateTodo is the function for an insert query on the db
func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?, ?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo is the function for a delete query on the db
func DeleteTodo(name string) error {
	deleteQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	defer deleteQ.Close()
	if err != nil {
		return err
	}
	return nil
}
