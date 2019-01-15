package main
import "fmt"

func main() {
	var by int
	var current_year int
	var banhammer int
	fmt.Println("Введите год своего рождения!")
	fmt.Scan(&by)
	current_year = 2019
	banhammer = current_year - by

	if banhammer < 18 {
		fmt.Println("Доступ запрещён!")
	} else {
		fmt.Println("Добро пожаловать! Наслаждайтесь!")
	}
}