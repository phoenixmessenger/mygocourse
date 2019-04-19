package main

import (
	"net/http"
	"testform/handlers"
)

func main() {
	http.HandleFunc("/",
		func(rw http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" {
				handlers.GetForm(rw, r)
			} else {
				handlers.PostForm(rw, r)
			}

		})
	http.ListenAndServe(":8080", nil)
}
