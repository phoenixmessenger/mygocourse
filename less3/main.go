package main

import "fmt"

func main() {
	
	// Вывести нечетные числа от 3 до 99
	
	var i, n int
	var simple bool

	for i = 2; i <= 100; i++  {
		simple = true
		for n = 2; n < i; n++ {
			if i % n == 0 {
				simple = false
				break
				} 	
			}
			if simple == true {
				fmt.Println(n)
		}
	}
}