package main

import (
	"log"
	"math/rand"
	"time"
)

func slice(a []int) int {
	rand.Seed(time.Now().UnixNano())

	a = make([]int, 10)
	var sum int
	for i := range a {
		a[i] = rand.Intn(100)
		sum += a[i]
	}
	log.Println(a)
	log.Println(sum)
	return sum
}

func main() {
	var b []int
	result := slice(b)
	avg := float64(result) / 10
	log.Println(avg)
}
