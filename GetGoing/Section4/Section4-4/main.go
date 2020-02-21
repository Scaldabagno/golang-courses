package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Received")
		fmt.Println(r.Method) // print method
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
