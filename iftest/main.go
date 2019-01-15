package main
import "fmt"

func main()  {
	d := -5

	if d > 0 {
		x1 := 8
		x2 := 16
		fmt.Printf("x1 = %v; x2 = %v\n", x1, x2) 
	} else if d == 0 {
		x := 19
		fmt.Printf("x = %v\n", x)
	} else {
		fmt.Println("No solutions!")
	}
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

		var mark int
		mark = 1
	 	if mark == 1 {
		 fmt.Println("Ужос!:(((((")
	 	} else if mark == 2 {
		 fmt.Println("Неуд!!!:////")
	 	} else if mark == 3 {
			fmt.Println("Удовлетворительно!")
		} else if mark == 4 {
			fmt.Println("Хорошо!")
		} else if mark == 5 {
			fmt.Println("Отлично! :)")
		} else {
			fmt.Println("Неверная оценка!")
		}

	}
	
