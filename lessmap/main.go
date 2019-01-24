package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	//map [Type]Type

	var m map[string]string

	m = make(map[string]string)
	m["name"] = "finalist"
	m["year"] = "1983"

	fmt.Println(m)

	//сделать map  интами, в котором slice с интами; создать десять элементов, с ключом К  и значением []int с рандомной длиной и рандомным значением.
	// при генерации следующего K проверить на наличие К в map. Если уже есть - сообщить и остановить цикл.

	var l map[int][]int
	var k, n, c int

	l = make(map[int][]int)
	for i := 0; i < 10; i++ {
		k = rand.Intn(20)
		v := l[k]

		if v == nil {
			n = rand.Intn(20) + 1
			v = make([]int, 0)
			for j := 0; j < n; j++ {
				c = rand.Intn(40)
				v = append(v, c)
			}
		}
		l[k] = v
	}
	fmt.Println(l)

	//map стран и слайз - список городов

	var b map[string][]string

	b = make(map[string][]string)
	slUkraine := make([]string, 0)
	slUkraine = append(slUkraine, "Kiev", "Niko", "Kherson")
	b["ua"] = slUkraine
	fmt.Println(b)

	slLatviia := make([]string, 0)
	slLatviia = append(slLatviia, "Riga", "Liepaja", "Jurmala")
	b["lv"] = slLatviia
	fmt.Println(b)

	/* */

	var z  = map[string]int {
		"test":rand.Intn(100),
		"mest":8,
		"dest":9,
		"jest":10,
	}
	fmt.Println(z)
	delete(z, "jest")
	fmt.Println(z)
	
	/* */
	var exist bool
	_,exist = z["jest"]
	if exist {
		panic("nojest")
	}

	/* */
	for key, value := range z {
		fmt.Printf("[%s] = [%d]\n", key, value)
	}
}
