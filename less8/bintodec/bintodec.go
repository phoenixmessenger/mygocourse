package dconv

import (
	"math"
	"strconv"
)

func BinToDec(s string) int {
	var cnt int
	cnt = len(s) - 1 //12

	if len(s) == 0 {
		return 0
	}

	//todo: check for 0 and 1 only
	var result int
	for _, e := range s {
		var currentNubmer string
		currentNubmer = string(e)
		number, _ := strconv.Atoi(currentNubmer)
		result += int(math.Pow(float64(number), float64(cnt)))
	}
	return result
}
