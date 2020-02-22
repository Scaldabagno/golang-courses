package controller

import (
	"net/http"
)

// Register endpoints
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	return mux
}
