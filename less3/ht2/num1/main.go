package main 
import "fmt"

func main() {
	// ввести четырехзначное число и вывести из него третью цифру

	var a, b int
	var res int
	
	fmt.Println("Введите четырёхзначное число!")
	fmt.Scan(&a)
	res = a % 10
	b = a % 100
	for ; res <= 9; res++ {
		if res != 0 {
			res = (a/10) % 10
		fmt.Printf("Искомое число равно %d\n", res)
		break;
		}
		if res == 0 {
			b = (a/10) % 10
			if b != 0 {
			fmt.Printf("Искомое число равно %d\n", b)
			break;	
			} else {
			fmt.Println("Искомое число равно 0!")
			break;
			}
		}
	}
}