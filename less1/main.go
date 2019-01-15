package main
import "fmt"

func main() {
	var bd int
	bd = 4
	bm := 7
	var pr int = 183
	pv := 98.9 //мой типа вес!!! в другой записи var pv float32 = 98.9
	var name string
	name = "Pavel"
	pv = pv + 2.1 //меню
	var a int
	a = 33 + 41 + 25 + 36
	srar := a/4;
	srfl := float32(a)/4

	fmt.Println(bd)
	fmt.Println(pr)
	fmt.Println(name)
	fmt.Println(bm)
	fmt.Printf("%.2f\n", pv)
	fmt.Print("abc\ncba\n")
	fmt.Printf("Your name is %s and your birthday is %d.%02d\n ", name, bd, bm)
	fmt.Println(srar)
	fmt.Println(srfl)
}
