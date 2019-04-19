package main

import "fmt"

func main() {

	var n int
	const symbol = "*"
	n = 5
	z := 1
	p := n - 1

	for i := 0; i < n; i++ {

		for ip := 0; ip < p; ip++ {
			fmt.Print(" ")
		}
		for iz := 0; iz < z; iz++ {
			fmt.Print(symbol)
		}
		fmt.Print("\n")
		p--
		z = z + 2
	}
	//fmt.Printf("p: %d, z: %d\n", p, z)
	p += 2
	z = z - 4
	for j := 0; j < n; j++ {
		for ijp := 0; ijp < p; ijp++ {
			fmt.Print(" ")
		}
		for ijz := 0; ijz < z; ijz++ {
			fmt.Print(symbol)
		}
		fmt.Print("\n")
		p++
		z = z - 2
	}
}
