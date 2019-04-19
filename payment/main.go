package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*Сгенерить map[string] int, где ключ - фамилия сотрудника, а значение - его зарплата.*/
	rand.Seed(time.Now().UnixNano())

	//генерируем
	var z = map[string]int{
		"Ivanov":    rand.Intn(1500),
		"Petrov":    rand.Intn(1500),
		"Sidorov":   rand.Intn(1500),
		"Berger":    rand.Intn(1500),
		"Goldstein": rand.Intn(1500),
	}
	fmt.Println(z)
	//исключение
	for key, value := range z {
		if value < 500 {
			delete(z, key)
			fmt.Printf("[%s] with zp[%d] removed!\n", key, value)
		} else if value > 1000 {
			delete(z, key)
			fmt.Printf("[%s] with zp[%d] removed!\n", key, value)
		}
	}
	fmt.Println(z)

	// Вариант решения Виталиком
	var names = []string{"s1", "s2", "s3", "s4", "s5"}
	var zp = make(map[string]int)
	for _, name := range names {
		zp[name] = rand.Intn(1500) + 1
	}
	fmt.Println(zp)
	for n, sum := range zp {
		if sum > 500 && sum < 1000 {
			continue
		}
		delete(zp, n)
		fmt.Printf("[%s] with zp[%d] removed!\n", n, sum)
	}
	fmt.Println(zp)
}
