package handlers

import (
	"fmt"
	"net/http"
)

func GetForm(rw http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(rw, "<h1>Представьтесь, пожалуйста!</h1>")
	fmt.Fprintf(rw, `<form action = "/" method = "post">
					<input type = "text" name = "name" placeholder = "Ваше имя">
					<button type = "submit">Отправить</button>
					</form>`)
}
