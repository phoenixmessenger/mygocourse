package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("Введите число a")
	fmt.Scan(&a)
	fmt.Println("Введите число b")
	fmt.Scan(&b)
	fmt.Println("Введите число c")
	fmt.Scan(&c)

	var D float64
	var disqrt float64
	D = math.Pow(b, 2) - 4*a*c
	disqrt = math.Sqrt(float64(D))

	if D > 0 {
		var x1, x2 float64
		x1 = float64(-b) + disqrt / 2*a
		x2 = float64(-b) - disqrt / 2*a
		fmt.Printf("Уравнения имеет два вещественных корня:\nx1 = %.2f\nx2 = %.2f\n", x1, x2)
	} else if D == 0 {
		var x float64
		x = float64(-b) / 2*a
		fmt.Printf("Уравнение имеет один корень: x = %.2f\n", x)
	} else {
		fmt.Println("Уравнение корней не имеет!")
	}
}