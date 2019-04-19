package main

import (
	"fmt"
	"math/rand"
	"time"
)

type AirReis struct {
	id    int
	from  string
	to    string
	Date  int
	Price int
}

type Customer struct {
	FirstName string
	LastName  string
	Ballance  int
}

var cities = []string{"Dublin", "Kiev", "Moskow", "SPB", "Odessa", "Kherson", "NY", "LasVegas", "London",
	"Tokyo", "Singapur", "Kualalumpur", "Sidney", "CapeTown", "Nayrobi", "Kasablanka"}

func buyReis(rid int, c Customer, Reises [][]AirReis) error {
	var customer Customer
	customer.FirstName = "John"
	customer.LastName = "Smith"
	customer.Ballance = 1935
	fmt.Print("Введите номер рейса для покупки!")
	fmt.Scan(&rid)

	//return err
}

func printReises(R [][]AirReis, limit int) {
	for _, e := range R {
		limit--
		if limit < 0 {
			break
		}
		if len(e) == 1 {
			fmt.Printf(`Без пересадок
Рейс с #%d
%s -> %s
Price: %d
`, e[0].id, e[0].from, e[0].to, e[0].Price)
		} else {
			fmt.Printf(`С пересадкой в %s:
Рейс c #%d
%s -> %s
Price: %d
===
Рейс c #%d
%s -> %s
Price: %d 
Total Price: %d
`, e[0].to, e[0].id, e[0].from, e[0].to, e[0].Price, e[1].id, e[1].from, e[1].to, e[1].Price, e[0].Price+e[1].Price)
		}
		fmt.Println("--------")
	}
}

func search(f, t string, df, dt int, reises []AirReis) [][]AirReis {
	var sorting [][]AirReis
	for _, e := range reises {
		if e.from == f && e.to == t {
			if e.Date >= df && e.Date <= dt {
				var tmp []AirReis
				tmp = append(tmp, e)
				sorting = append(sorting, tmp)
			}
		}
	}

	var fly1, fly2 []AirReis
	for _, e := range reises {
		if e.from == f &&
			e.to != t &&
			e.Date >= df &&
			e.Date <= dt {
			fly1 = append(fly1, e)
		}
	}
	//log.Println(len(fly1), fly1, fly2)
	for _, e := range reises {
		if e.Date >= df &&
			e.Date <= dt &&
			e.from != f &&
			e.to == t {
			fly2 = append(fly2, e)
		}
	}
	//log.Println(len(fly2), fly2)
	// var speresadkoy [][]AirReis
	for _, f1 := range fly1 {
		for _, f2 := range fly2 {
			if f1.to == f2.from {
				var f1f2 []AirReis
				f1f2 = append(f1f2, f1)
				f1f2 = append(f1f2, f2)
				sorting = append(sorting, f1f2)
			}
		}
	}
	//log.Println(len(speresadkoy), speresadkoy)

	// var spmin, index int
	// spmin = speresadkoy[0][0].Price + speresadkoy[0][1].Price

	// for in, e := range speresadkoy {
	// 	sum := e[0].Price + e[1].Price
	// 	if spmin > sum {
	// 		spmin = sum
	// 		index = in
	// 	}
	// }
	// log.Println(speresadkoy[index], spmin)

	return sorting
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const reiscount = 15000
	var reises = make([]AirReis, 0, reiscount)
	for i := 0; i < reiscount; i++ {
		var reis AirReis
		reis.id = i + 1
		reis.from = cities[rand.Intn(len(cities))]
		reis.to = cities[rand.Intn(len(cities))]
		reis.Date = rand.Intn(31) + 1
		reis.Price = rand.Intn(2000) + 1000

		reises = append(reises, reis)
	}
	var f, t string
	f = "Dublin"
	t = "CapeTown"
	df := 3
	dt := 3
	sorting := search(f, t, df, dt, reises)
	printReises(sorting, 6)

	var customer Customer
	customer.FirstName = "John"
	customer.LastName = "Smith"
	customer.Ballance = 1935
	fmt.Print("Введите номер рейса для покупки!")
	var rid int
	fmt.Scan(&rid)
}
