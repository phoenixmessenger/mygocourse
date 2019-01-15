package main

import "fmt" 

func main() {
	
	// палиндром

	var s string
	fmt.Println("Введите слово или предложение!")
	fmt.Scan(&s)
	n := len(s) - 1 
	for e := range s {	
		if s[e] == s[n] {
			fmt.Println("Это  палиндром!")
			return
		} else {
			n--
			fmt.Println("Это не палиндром!")
		}
	}
}