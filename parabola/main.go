package main

import "fmt"

func main() {
	
	var x int
	var xm []int
	var ym []int
	xm = make([]int, 200)
	ym = make([]int, 200)
	for i := range xm {
		for x = -100; x < i-99; x++ {
			xm[i] = x
			ym[i] = x * x
		}
	}
	fmt.Println(xm)
	fmt.Println(ym)
}
