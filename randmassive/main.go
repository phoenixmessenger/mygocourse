package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var min, max int
	var avg float64
	var arr [10]int

	for i := range arr {
		arr[i] = rand.Intn(100)
		max = arr[0]
	}
	for _, e := range arr {
		if max < e {
			max = e
		} else if min > e {
			fmt.Println(min)
		}
		avg = float64(avg) + float64(e)
		avg = float64(avg) / 10
	}
	fmt.Printf("%d\n", arr)
	fmt.Printf("avg: %.3f\nmin:%d\nmax:%d\n", avg, min, max)
}
