package dconv

import "strconv"

var hexnumbers map[int]string

func DecToHex(n int) string {
	var reshex string
	var ostatok int

	for n > 15 {
		ostatok = n % 16
		if ostatok > 9 {
			reshex += hexnumbers[ostatok]

		} else {
			reshex += strconv.Itoa(ostatok)
		}
		n /= 16
	}
	if n > 9 {
		reshex += hexnumbers[n]
	} else {
		reshex += strconv.Itoa(n)
	}
	return turnMass(reshex)
}

func init() {
	hexnumbers = make(map[int]string)
	hexnumbers[10] = "A"
	hexnumbers[11] = "B"
	hexnumbers[12] = "C"
	hexnumbers[13] = "D"
	hexnumbers[14] = "E"
	hexnumbers[15] = "F"
}
