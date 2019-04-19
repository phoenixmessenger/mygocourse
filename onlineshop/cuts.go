package main

import "log"

func createCats() []Category {
	var names = []string{"Хлеб", "Выпечка"}
	var id int = 1

	var res []Category

	for _, name := range names {
		var c Category
		c.ID = id
		c.Title = name
		id++
		res = append(res, c)
	}
	return res
}

func printCats(c []Category) {
	for _, e := range c {
		log.Printf("%d) %s\n", e.ID, e.Title)
	}
}
