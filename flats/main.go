package main
import "fmt"

func main() {
	var flat int
	fmt.Println("Введите номер квартиры")
	fmt.Scan(&flat)

	if flat < 20 && flat > 0 {
		if flat >= 1 && flat <= 4 {
		fmt.Println("Находится на первом этаже первого подъезда!")
	}  else if flat >= 5 && flat <= 8 {
		fmt.Println("Находится на втором этаже первого подъезда!")
	} else if flat >= 9 && flat <= 12 { 
		fmt.Println("Находится на третьем этаже первого подъезда!")
	} else if flat >= 13 && flat <= 16 { 
		fmt.Println("Находится на четвертом этаже первого подъезда!")
	} else if flat >= 17 && flat <= 20 { 
		fmt.Println("Находится на пятом этаже первого подъезда!")
		}
	
	} else if flat > 20 && flat < 40 {
		if flat >= 21 && flat <= 24 {
			fmt.Println("Находится на первом этаже второго подъезда!")
		}  else if flat >= 25 && flat <= 28 {
			fmt.Println("Находится на втором этаже второго подъезда!")
		} else if flat >= 29 && flat <= 32 { 
			fmt.Println("Находится на третьем этаже второго подъезда!")
		} else if flat >= 33 && flat <= 36 { 
			fmt.Println("Находится на четвертом этаже второго подъезда!")
		} else if flat >= 37 && flat <= 40 { 
			fmt.Println("Находится на пятом этаже второго подъезда!")
			}

	} else if flat > 40 && flat < 60 {
		if flat >= 41 && flat <= 44 {
			fmt.Println("Находится на первом этаже третьего подъезда!")
		}  else if flat >= 45 && flat <= 48 {
			fmt.Println("Находится на втором этаже третьего подъезда!")
		} else if flat >= 49 && flat <= 52 { 
			fmt.Println("Находится на третьем этаже третьего подъезда!")
		} else if flat >= 53 && flat <= 56 { 
			fmt.Println("Находится на четвертом этаже третьего подъезда!")
		} else if flat >= 57 && flat <= 60 { 
			fmt.Println("Находится на пятом этаже третьего подъезда!")
			}

	} else if flat > 60 && flat < 80 {
		if flat >= 61 && flat <= 64 {
			fmt.Println("Находится на первом этаже четвертого подъезда!")
		}  else if flat >= 65 && flat <= 68 {
			fmt.Println("Находится на втором этаже четвертого подъезда!")
		} else if flat >= 69 && flat <= 72 { 
			fmt.Println("Находится на третьем этаже четвертого подъезда!")
		} else if flat >= 73 && flat <= 76 { 
			fmt.Println("Находится на четвертом этаже четвертого подъезда!")
		} else if flat >= 77 && flat <= 80 { 
			fmt.Println("Находится на пятом этаже четвертого подъезда!")
			}
		
	} else if flat <= 0 {
		fmt.Println("Неверный формат ввода!")

	} else {
		fmt.Println("Нет такой квартиры!")
	}
}
