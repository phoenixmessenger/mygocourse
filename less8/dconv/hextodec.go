package dconv

import "math"

var HexToInt = map[string]int{
	"F": 15,
	"E": 14,
	"D": 13,
	"C": 12,
	"B": 11,
	"A": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"1": 1,
	"0": 0,
}

func HexToDec(s string) int {
	var n = len(s) - 1
	var result int

	for _, e := range s {
		tmp := HexToInt[string(e)]
		tmp *= int(math.Pow(16.0, float64(n)))
		result += tmp
		n--
	}
	return result
}
