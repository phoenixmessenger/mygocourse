package main

import (
	"log"
	"math/rand"
	"time"
)

func shuffleBox(box []Sock) []Sock {
	rand.Seed(time.Now().UnixNano())
	var mix []Sock
	mix = make([]Sock, 0, len(box))

	for len(box) > 0 {
		rnd := rand.Intn(len(box))
		mix = append(mix, box[rnd])
		box = append(box[:rnd], box[rnd+1:]...)
	}
	return mix
}

func main() {
	var box []Sock
	for i := 0; i < 7; i++ {
		box = append(box, BLACK)
		box = append(box, WHITE)
	}

	for i := 0; i < 8; i++ {
		box = append(box, RED)
		box = append(box, GREEN)
	}

	log.Println(box)
	box = shuffleBox(box)
	log.Println(box)

	var status bool

	var sortSock = make([]Sock, 0)
	sortSock = append(sortSock, box[0])
	for n := 1; n < len(box); n++ {
		for _, e := range sortSock {
			if e != box[n] {
				continue
			}
			status = true
			break
		}
		if status == true {
			log.Println(n + 1)
			break
		}
		sortSock = append(sortSock, box[n])
	}
}
