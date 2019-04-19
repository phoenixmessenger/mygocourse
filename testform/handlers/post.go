package handlers

import (
	"fmt"
	"net/http"
)

func PostForm(rw http.ResponseWriter,
	r *http.Request) {
	fmt.Println("received post request")
	fmt.Printf("ip: %s", r.RemoteAddr)
	var name = r.PostFormValue("name")
	fmt.Printf("name: %s\n", name)
	fmt.Fprintf(rw, "<h1>Привет, <i>%s\n</i></h1>", name)

}
