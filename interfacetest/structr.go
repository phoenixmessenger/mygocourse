package main

import "fmt"

type User struct {
	Email    string
	Password string
	Name     string
}

type Product struct {
	Art   string
	Name  string
	Price int64
}

type Cart struct {
	User    *User
	Product []*Product
}

func PrintData(i interface{}) {

	fUser, ok := i.(User)
	if ok {
		fmt.Printf(`
		Email: %v
		Password: %v
		Name: %v
		`, fUser.Email, fUser.Password, fUser.Name)
		return
	}

	fProduct, ok := i.(Product)
	if ok {
		fmt.Printf(`
		Art: %v
		Name: %v
		Price: %v
		`, fProduct.Art, fProduct.Name, fProduct.Price)
		return
	}

	fCart, ok := i.(*Cart)
	if ok {
		PrintData(fCart.User)
		for _, e := range fCart.Product {
			PrintData(e)
		}
		return
	}
}

func main() {
	u := &User{"test@test.com", "abc123", "Fedor"}
	p := &Product{"Notebook", "Lenovo", 1000}
	PrintData(*u)
	PrintData(*p)

	cart := &Cart{
		&User{"test@test.com", "abc123", "Fedor"},
		[]*Product{
			&Product{"Notebook", "Lenovo", 1000},
			&Product{"Notebook", "Samsung", 2000},
			p,
		},
	}
	//append
	cart.Cart = append()
}

//д.з дбавить структуры Library и Shelf
// //type Library struct {
// 	Name
// 	Address
// 	[]Shelf
// }
