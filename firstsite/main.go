package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "<h1>Hello <i>%s</i></h1>", "finalist")
	})

	http.ListenAndServe(":8080", nil)
}
