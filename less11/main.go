package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func IsPalindrome(word string) bool {
	if utf8.RuneCountInString(word) <= 1 {
		return true
	}
	first, sizeOfFirst := utf8.DecodeRuneInString(word)
	last, sizeOfLast := utf8.DecodeLastRuneInString(word)
	if first != last {
		return false
	}
	return IsPalindrome(word[sizeOfFirst : len(word)-sizeOfLast])
}

func main() {
	var pal string
	fmt.Print("Enter your word!")
	fmt.Scanf("%s", &pal)
	pal1 := IsPalindrome(pal)
	log.Println(pal1)
}
