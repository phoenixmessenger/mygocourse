package main

import (
	"fmt"
	"log"
)

// Организовать беспрерывный ввод чисел с клавиатуры, пока пользователь не введёт 0. После ввода нуля,
// показать на экран количество чисел, которые были введены, их общую сумму и среднее арифметическое

func main() {
	var enternum int
	var cnum, sum int
	var avg float64
	fmt.Printf("Введите любое целое число\n")

	for {
		fmt.Scan(&enternum)
		if enternum == 0 {
			break
		}
		cnum++
		sum += enternum
		avg = float64(sum) / float64(cnum)
	}
	log.Println(cnum)
	log.Println(sum)
	log.Println(avg)
}
