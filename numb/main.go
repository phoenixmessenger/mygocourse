package main
import "fmt"

func main() {
var n1, n2, n3 int
n1 = -3
n2 = -8
n3 = -16

//счетчик
var count int
count = 0
if n1 > 0 {
	count++
} 
if n2 > 0 {
	count++
} 
if n3 > 0 {
	count++
} 

fmt.Printf("%d pozitive numbers!\n", count)

//с операторами 
if n1 > 0 && n2 > 0 && n3 > 0 {
	fmt.Println("Three pozitive numbers")
} else if n1 > 0 && n2 > 0 && n3 < 0 {
	fmt.Println("Two pozitive numbers")
} else if n1 > 0 && n2 < 0 && n3 < 0 {
	fmt.Println("One pozitive number")
} else {
	fmt.Println("No pozitive numbers")
}
}