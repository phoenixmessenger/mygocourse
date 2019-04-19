package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

// Есть четыре созвездия, много звезд в них и нужно
// сгенерировать их на трехмерной плоскости.
// Сделать функцию, которая определяет дальность расстояния между звездами
// (самое короткое и самое длинное расстояния)

type Star struct {
	name string
	x    int
	y    int
	z    int
}

type Galaxy struct {
	Name string
	star []Star
}

func main() {
	var raw_stars = map[string][]string{
		"Андромеда": {"Веритате", "Нембус", "Альферац", "Мирах", "Аламах", "Адхил", "Титавин"},
		"Близнецы":  {"Кастор", "Поллукс", "Альхена", "Васат", "Мебсута", "Мекбуда", "Пропус"},
		"Большая Медведица": {"Интеркрус", "Чалаван", "Алькор", "Дубхе", "Мерак", "Фекда", "Мегрец",
			"Алиот", "Мицар", "Алкаид", "Талитха"},
		"Большой Пёс": {"Сириус", "Мирзам", "Мулифен", "Везен", "Адара", "Фуруд", "Алудра",
			"Исида", "Унургуните"},
	}
	var universe []Galaxy

	for g, strs := range raw_stars {
		galaxy := Galaxy{}
		galaxy.Name = g

		for _, s := range strs {
			star := Star{}
			star.name = s
			star.x = rand.Intn(1000000)
			star.y = rand.Intn(5000000)
			star.z = rand.Intn(9000000)
			galaxy.star = append(galaxy.star, star)
		}
		universe = append(universe, galaxy)
	}
	PrintUnverse(universe)

	var max, min float64
	var maxStar1, maxStar2 string
	var minStar1, minStar2 string

	for _, galaxy1 := range universe {
		for _, star1 := range galaxy1.star {
			for _, galaxy2 := range universe {
				for _, star2 := range galaxy2.star {
					if star1.name == star2.name {
						continue
					}
					dist := Distance(star1, star2)
					if max < dist {
						max = dist
						maxStar1 = star1.name
						maxStar2 = star2.name
					}
					if min == 0 {
						min = dist
					} else {
						if min > dist {
							min = dist
							minStar1 = star1.name
							minStar2 = star2.name
						}
					}
				}
			}
		}
	}
	log.Printf("Максимальное расстояние между звездами %s и %s: %v\n", maxStar1, maxStar2, max)
	log.Printf("Минимальное расстояние между звездами %s и %s: %v\n", minStar1, minStar2, min)

}

func PrintUnverse(u []Galaxy) {
	for _, g := range u {
		fmt.Printf("Galaxy: %s\n", g.Name)
		for _, s := range g.star {
			fmt.Printf("\tStar: %s; %d; %d; %d\n", s.name, s.x, s.y, s.z)
		}
	}
}

func Distance(a Star, b Star) float64 {
	var c int
	c = ((a.x - b.x) * (a.x - b.x)) + ((a.y - b.y) * (a.y - b.y)) + ((a.z - b.z) * (a.z - b.z))
	return math.Sqrt(float64(c))
}
