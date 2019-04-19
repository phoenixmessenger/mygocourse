package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scan(&n)
	fmt.Printf("multitable for %d ", n)

	for i := 1; i < 11; i++ {
		fmt.Printf("%d * %d = %d\n", i, n, i*n)
	}
}
