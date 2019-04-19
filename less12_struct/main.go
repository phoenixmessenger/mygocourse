package main

import (
	"log"
)

type Car struct {
	wheels int
	doors  int
	engine float32
}

func main() {
	//описать 10k авто
	var c1 Car
	var c10k [3]Car
	c1.wheels = 4
	c1.doors = 4
	c1.engine = 1.8
	log.Println(c1)
	log.Printf("%+v", c1)
	log.Printf("%+v", c10k)

	//квадратное уравнение
	var u KvUravn
	u.a = 1
	u.b = -6
	u.c = 9
	res := Solution(u)
	log.Println(res)
}
