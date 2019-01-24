package main

import (
	"log"
	"math/rand"
	"time"
	"errors"

)

func convertUAHtoUSD(sum int) (int, error) {
	rand.Seed(time.Now().UnixNano())

	const curs = 27
	var vero = rand.Intn(100)
	if vero < 50 {
		errors.New("server error!")


	}
	var res int
	res = sum / curs
	return res, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var s, res int
	var err error
	s = 9141
	res, err = convertUAHtoUSD(s)
	if err != nil {
		log.Printf("convert error: %s", err)
		return
	}
	log.Printf("%d uah = %d usd!\n", s, res)
}
