package main

import (
	"log"
	"math/rand"
)

//дан массив из 150 рандомных чисел.
//Если записываемое число больше предыдущего -
// делим на джва

func main() {
	const n = 150
	var numbs = make([]int, n)
	for i := 0; i < n; i++ {
		a := rand.Intn(1500)
		if i == 0 {
			numbs[i] = a
		} else {
			if numbs[i] > numbs[i-1] {
				a /= 2
			}
			numbs[i] = a
		}
	}
	log.Println(numbs)
}
