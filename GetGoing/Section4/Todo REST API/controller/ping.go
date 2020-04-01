package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"../views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			log.Println("-> ping()")
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
			log.Println("<- ping()")
		}
	}
}
