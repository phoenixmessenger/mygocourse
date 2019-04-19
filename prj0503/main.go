package main

import "fmt"

type SortingInterface interface {
	Len() int
	Less(a, b int) bool
	Swap(a, b int)
}

func Sort(data SortingInterface) {
	for i := 0; i < data.Len()-1; i++ {
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(i, j) {
				data.Swap(i, j)
			}
		}
	}
}

type StringArray struct {
	Strings []string
}

type FloatArray struct {
	Float []float64
}

func (s *StringArray) Len() int {
	return len(s.Strings)
}

func (s *StringArray) Less(a, b int) bool { //узнать, меньше ли єлемент под номером а єлемента под номером b?
	if s.Strings[a] < s.Strings[b] {
		return true
	} else {
		return false
	}
}

func (s *StringArray) Swap(a, b int) {
	s.Strings[a], s.Strings[b] = s.Strings[b], s.Strings[a]
}

func (s *StringArray) Print() {
	for _, e := range s.Strings {
		fmt.Println(e)
	}
	fmt.Println("-------")
}

func (f64 *FloatArray) Len() int {
	return len(f64.Float)
}

func (f64 *FloatArray) Less(a, b int) bool {
	return f64.Float[a] < f64.Float[b]
}

func (f64 *FloatArray) Swap(a, b int) {
	f64.Float[a], f64.Float[b] = f64.Float[b], f64.Float[a]
}

func (f64 *FloatArray) Print() {
	for _, e := range f64.Float {
		fmt.Println(e)
	}
	fmt.Println("-------")
}

//
type Book struct {
	Author string
	Pages  int
	Title  string
	Price  int
}

type Library struct {
	Books     []*Book
	SortType  string
	Direction string
}

func (self *Library) Len() int {
	return len(self.Books)
}

func (self *Library) Swap(i, j int) {
	self.Books[i], self.Books[j] = self.Books[j], self.Books[i]
}

func (self *Library) Less(i, j int) bool {
	switch self.SortType {
	case "Author":
		if self.Direction == "asc" {
			return self.Books[i].Author > self.Books[j].Author
		} else {
			return self.Books[i].Author < self.Books[j].Author
		}
	case "Pages":
		if self.Direction == "asc" {
			return self.Books[i].Pages > self.Books[j].Pages
		} else {
			return self.Books[i].Pages < self.Books[j].Pages
		}
	case "Title":
		if self.Direction == "asc" {
			return self.Books[i].Title > self.Books[j].Title
		} else {
			return self.Books[i].Title < self.Books[j].Title
		}
	case "Price":
		if self.Direction == "asc" {
			return self.Books[i].Price > self.Books[j].Price
		} else {
			return self.Books[i].Price < self.Books[j].Price
		}
	}
	return false
}

func (self *Library) Print() {
	for _, e := range self.Books {
		fmt.Println(e)
	}
	fmt.Printf("--------\n")
}

func main() {
	var s StringArray
	s.Strings = []string{"qweely", "ABC", "repchik", "Eyjafjallajökull", "mafioznik"}
	s.Print()
	Sort(&s)
	s.Print()

	var f FloatArray
	f.Float = []float64{2.5, 3.6, 4.7, 5.8, 1.2}
	f.Print()
	Sort(&f)
	f.Print()

	var lbr = Library{
		Books: []*Book{
			&Book{"Andjei Sapkowsky", 300, "The Witcher. Elves blood", 300},
			&Book{"Joahnn Rowling", 672, "Harry Potter and the Goblet of Fire", 160},
			&Book{"J.R.R. Tolkien", 252, "The Lord of the Rings. Pt.1: The fellowship of the Ring", 520},
		},
	}

	lbr.SortType = "Price"
	lbr.Direction = "asc"
	lbr.Print()
	Sort(&lbr)
	lbr.Print()

}
