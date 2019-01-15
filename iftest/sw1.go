package main
import "fmt"

func main() {
	a := 0
		switch a {
			case 1: 
				fmt.Println("1")
			case 2: 
				fmt.Println("2")
			case 3: 
				fmt.Println("3")
			case 4, 5:
				fmt.Println("4 or 5")
			default:
				fmt.Println("Other!")
			}
	}