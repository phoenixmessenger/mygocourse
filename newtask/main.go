package main
import ("fmt" 
"math/rand"
"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	/*создать четыре массива инта*/
	var mat []int
	var phis []int
	var penie []int
	var fizra []int 
	var avgmat float64
	var sum int
	var result(map[string] float64)
	mat = make([]int, 20)
	phis = make([]int, 20)
	penie = make([]int, 20)
	fizra = make([]int, 20)
	result = make(map[string] float64)

	for i := 0; i < 20; i++ {
		mat[i] = rand.Intn(12)
		for _,e := range mat {
			sum += e

		}
	}
	for a := range phis {
		phis[a] = rand.Intn(12)
	}
	for b := range penie {
		penie[b] = rand.Intn(12)
	}
	for c := range fizra {
		fizra[c] = rand.Intn(12)
	}
	fmt.Println(mat)
	fmt.Println(phis)
	fmt.Println(penie)
	fmt.Println(fizra)


	avgmat = float64(sum) / float64(20)
	result ["math"] = avgmat
	fmt.Println(result)
}