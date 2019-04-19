package main

import "fmt"

func main() {
	cats := createCats()
	bakes := createBake()
	var yourcar []Bake
	var choice int
	for {
		mainMenu()
		fmt.Scan(&choice)
		switch choice {
		case 0:
			return
		case 1:
			printCats(cats)
		case 2:
			fmt.Print("Выберите категорию!")
			fmt.Scan(&choice)
			printBakesInCat(choice, bakes)
		case 3:
			fmt.Print("Введите ID товара!")
			fmt.Scan(&choice)
			AddToCar(choice, bakes)
		case 4:
			fmt.Print("Список выбранных вами товаров!")
			AddToCar(choice, yourcar)
		}
	}
}

func mainMenu() {
	fmt.Println("-----------")
	fmt.Println("1 - Список категорий")
	fmt.Println("2 - Список выпечки")
	fmt.Println("3 - Купить")
	fmt.Println("4 - Добавить в корзину")
	fmt.Println("0 - Выход")
}
