package main

import ("fmt"
		"time"
		"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	const n = 9
	var arr = [ ]int {
		1,8,5,3,4,9999,7,18,100}

	for i, e := range arr {   //range - граница, длина границы. 
		fmt.Printf("Array[%d] = %d\n", i, e) //Проходя по всей длине массива arr  - получаем индекс i в значении е 
	}

var a[10] int
var sum int 
	for i := range a {
		a[i] = rand.Int()
	}
	for _,e := range a {
		sum += e
	}
	
for v1 := range a {  //v1-index
	sum += v1
	}
for v1, v2 := range a {  //v1,-index v2,-index
	sum += v1
	sum += v2
	}
	fmt.Println(sum)	

	// отзеркаливаем целочисленные элементы массива
	
	var arr2[10] int
	var reverse[10] int
	var i int
	for i = range arr2 {
		arr2[i] = rand.Intn(10)

	}
	fmt.Println(arr2)
	
	var r int 
	r = 9
	for i = range reverse {
		reverse[r] = arr2[i]
		r--
		
	}
	fmt.Println(reverse)
}