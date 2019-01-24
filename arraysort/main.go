package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var arr [10]int
	var kv1, perc, kv3 int
	var avg float64
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	for i := 0; i < 10; i++ {

		for j := 0; j < 9; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
	kv1 = arr[2]
	perc = arr[5]
	kv3 = arr[7]
	for _, e := range arr {
		avg = float64(avg) + float64(e)
		avg = float64(avg) / 10
	}
	fmt.Printf("avg: %.3f\nkv1: %d\nperc: %d\nkv3: %d\n", avg, kv1, perc, kv3)

	arr2 := [...]int{1, 3, 8, 9, 10} /*встроенная функция appeng()*/
	fmt.Println(arr2)

	//	append([]Type,...Type)[]Type
	narr2 := append(arr2[0:5], 17)
	lastTwo := arr2[3:]
	threeEl := arr2[1:5]
	begin := arr2[:3]
	fmt.Println(lastTwo)
	fmt.Println(threeEl)
	fmt.Println(begin)
	fmt.Println(narr2)

	//создаем новый пустой массив:
	var slice []int
	narr2 = append(slice, 1) //slice + 1
	slice = append(slice, 3, 2, 8)
	slice[0] = 1
	slice = append(slice, 3)
	fmt.Println(len(slice))

	//make([]Type,... Type)[]Type;      make([]Type, int, int) - встроенная функция make создаёт новый срез
	var slice2 []int
	var slice3 []int
	slice2 = make([]int, 5)
	slice2[0] = 3
	slice2 = append(slice2, 10)
	fmt.Println(slice2)
	
	slice3 = make([]int, 0, 10)
	slice3 = append(slice3, 10,20,30,40,41,42,43,44,45,46,47,48,49,50)

	fmt.Println(slice3)

}