package controller

import (
	"net/http"
)

// Register endpoints
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/", crud())
	return mux
}
