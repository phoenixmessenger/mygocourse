package main
import "fmt"

func main() {
	var flat, level int
	fmt.Println("Введите номер квартиры")
	fmt.Scan(&flat)
	var number_level int
	var number_porch int
	level = 20 % flat
	number_porch = flat / 20 + 1

		if level != 0 {
			number_level = flat / 4 + 1
			fmt.Printf("Квартира находится на %d этаже %d подъезда\n", number_level, number_porch)
		} else if level == 0 {
			number_level = flat / 4
			fmt.Printf("Квартира находится на %d этаже %d подъезда\n", number_level, number_porch)			
		} else if flat > 80 {
		fmt.Println("Нет такой квартиры!")
	} else if flat <= 0{
		fmt.Println("Неверный формат ввода")
	}
}