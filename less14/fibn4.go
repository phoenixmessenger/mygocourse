package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var slays []int
	slays = make([]int, 20)
	slays[0] = rand.Intn(5)
	slays[1] = rand.Intn(20)
	for i := 2; i < 20; i++ {
		slays[i] = slays[i-2] + slays[i-1]
	}

	fmt.Println(slays)

}
