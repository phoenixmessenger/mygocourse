package main
import (
	"log"
	"math"
)
func distance1D(a, b int) int {
	d := a - b
	if d < 0 {
		d *= -1
	}
	return d
}
func distance2D(x1, x2, y1, y2 int) float64 {
	var dist int
	dist = ((x1 - x2) * (x1 - x2)) + ((y1 - y2) * (y1 - y2))
	var res float64
	res = math.Sqrt(float64(dist))
	return res
}
func main() {
	res1d := distance1D(6, 12)
	log.Println(res1d)
	res := distance2D(0, 3, 0, 4)
	log.Println(res)
}