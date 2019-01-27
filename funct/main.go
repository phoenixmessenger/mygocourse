package main
import "log"

func stepen(a, b int) int {
	
	var result int
	result = 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}


func main() {
	r := stepen(3, 8)
	log.Println(r)
}
