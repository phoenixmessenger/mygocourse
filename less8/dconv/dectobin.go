package dconv
import "strconv"

func DecToBin(d int) string {
	var res string
	for d > 0 {
		o := d % 2
		d = d/2
		res += strconv.Itoa(o)
	}
	return turnMass(res)
}
