package main 
import ( 
	"time"
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//Угадываем мы
	var n int
	var u int 
	n = rand.Intn(100)
	fmt.Println("Комп загадал число от 1 до 100, угадай!")

	for n != u {

		fmt.Print("Введи число: ")
		fmt.Scan(&u)
		if u < n {
			println("Больше!")
		} else {
			println("Меньше!")
		}
	}
	fmt.Println("Угадал!")
	
	//Угадывает компьютер
	
	var number int
	var ch string
	println("Введите число для угадывания!")
	fmt.Scan(&number)
	fmt.Println(number)

	for number = 1; number < 100; {
		var result int
		result = 50
		fmt.Println(result)
		
		for result = 1; result < 100; result++  {
			fmt.Scan(&ch)
			if ch == "<" {
				result = 100 - number
				fmt.Println(result)
			} else if ch == ">" {
				result =  result + 1
				fmt.Println(result)
			} else if ch == "=" {
				result = number
				fmt.Println("Угадал!")
				break
			}
		}
	break
	}

}		