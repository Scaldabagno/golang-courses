package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./structs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	address := "localhost:3000"
	fmt.Println("Server listening on " + address)
	http.ListenAndServe(address, mux)
}
