package controller

import (
	"net/http"

	_ "../model" // vndjisun
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// if error occurs
			if err := model.createTodo(); err != nil {
				w.Write([]byte("Some Error"))
				return
			}
		}
	}
}
