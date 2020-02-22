package main

import (
	"fmt"
	"net/http"

	"./controller"
)

func main() {
	mux := controller.Register()
	address := "localhost:3000"
	fmt.Println("Server listening on " + address)
	http.ListenAndServe(address, mux)
}
