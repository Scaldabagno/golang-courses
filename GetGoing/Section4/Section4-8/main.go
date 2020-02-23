package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // mysql driver

	"./controller"
	"./model"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	address := "localhost:3000"
	fmt.Println("Server listening on " + address)
	http.ListenAndServe(address, mux)
}
