/* Сортировка массивов пузырьками, но не со статическим, а с динамическим массивами: создаем срез, заполняем рандомом через append
сортируем пузырьком*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// объявляем slc и создаем прототип
	var slc []int
	slc = make([]int, 10)

	//заполняем рандомом и расширяем границы среза
	for i := range slc {
		slc[i] = rand.Intn(200)
		slc = append(slc, rand.Intn(200))
	}
	fmt.Println(slc)

	// сортируем методов пузырьков (создаём переменную, и с помощью её перезаписываем элементы массива)
	for i := 0; i < 20; i++ {
		for n := 0; n < 19; n++ {
			if slc[n] > slc[n+1] {
				slc[n], slc[n+1] = slc[n+1], slc[n]
			}
		}
	}
	fmt.Println(slc)
}
