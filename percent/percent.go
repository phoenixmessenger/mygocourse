package main

import "fmt"

func main() {

	var num int
	fmt.Printf("Введите целое число от 1 до 65535\n")
	fmt.Scan(&num)

	p10 := float32(num) / 100 * 10
	fmt.Printf("Десять процентов от вашего числа - это %.2f\n", p10)

	p15 := float32(num) / 100 * 15
	fmt.Printf("Пятнадцать процентов от вашего числа - это %.2f\n", p15)

	p25 := float32(num) / 100 * 25
	fmt.Printf("Двадцать пять процентов от вашего числа - это %.2f\n", p25)

}
