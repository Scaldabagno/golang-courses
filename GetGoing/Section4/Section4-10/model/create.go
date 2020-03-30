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
