package main

import (
	"fmt"
	"math/rand"
)

func createBake() []Bake {
	var names = []string{"Батон", "Бородинский", "Кирпич", "Элит", "Горожанин", "Зерновой",
		"Ватрушка", "Роллини", "Маковка", "Пончик", "Ромовая баба"}
	var id = 1
	var res []Bake

	for _, name := range names {
		var b Bake
		b.ID = id
		b.Title = name
		b.Price = rand.Intn(50) + 10
		b.Category = rand.Intn(2) + 1
		id++
		res = append(res, b)
	}
	return res
}
func printBakes(b []Bake) {
	for _, e := range b {
		fmt.Printf("%d) %s стоит %d грн.\n", e.ID, e.Title, e.Price)
	}
}

func printBakesInCat(cat int, bakes []Bake) {
	for _, e := range bakes {
		if e.Category != cat {
			continue
		} else {
			fmt.Printf("%d) %s стоит %d грн.\n", e.ID, e.Title, e.Price)
		}
	}
}

func buyBake(id int, bakes []Bake) {
	for _, e := range bakes {
		if e.ID != id {
			continue
		}
		fmt.Printf("Вы купили %s за %d грн\n", e.Title, e.Price)
		fmt.Printf("Приятного аппетита!")
	}
}

func AddToCar(id int, bakes []Bake) {
	for _, e := range bakes {
		if e.ID != id {
			fmt.Printf("Товар %s ценой в %d гривен добавлен в корзину\n", e.Title, e.Price)
			fmt.Printf("Выберите новый товар или нажмите 0 для выхода и подсчёта стоимости\n")
		}
	}
}
