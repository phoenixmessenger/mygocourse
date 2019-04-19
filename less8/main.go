package main

import (
	"less8/dconv"
	"log"
)

func main() {
	str := dconv.DecToBin(250)
	log.Println(str)
	v := "英文摘要"
	ba := []rune(v)
	for _, e := range ba {
		log.Printf("%d = %c!\n", e, e)
	}

	result := dconv.BinToDec(str)
	log.Println(result)

	resulthex := dconv.DecToHex(65535)
	log.Println(resulthex)

	resulthtod := dconv.HexToDec("FACB6D")
	log.Println(resulthtod)
}
