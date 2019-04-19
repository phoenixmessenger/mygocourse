// Д.з. создать проект. Массив строковых данных и отсортировать его.
// По этому образцу.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Student struct {
	name  string
	year  int
	group string
	marks []int
}

//гетеры и сетеры
func (s *Student) SetName(n string) {
	s.name = n
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) AddMark(m int) {
	s.marks = append(s.marks, m)
}

func (s *Student) Average() float64 {
	var avg float64
	for _, e := range s.marks {
		avg += float64(e)
	}
	avg /= float64(len(s.marks))
	return avg
}

func generateStudents(n int) []*Student {
	var s = make([]*Student, 0, n)
	var m = 5
	for n > 0 {
		n--
		tmp := new(Student)
		tmp.name = "Vitalii" + strconv.Itoa(n)
		tmp.year = rand.Intn(30) + 1980
		tmp.group = "Base-1 Go"
		for m = 0; m < 5; m++ {
			tmp.AddMark(rand.Intn(5) + 1)
		}
		s = append(s, tmp)
	}
	return s
}

func printStudens(s []*Student) {
	for _, e := range s {
		fmt.Printf("%+v\t%v\n", e, e.Average())
	}
}

func main() {
	// var s Student
	// s.SetName("Alyoshka")
	// s.AddMark(2)
	// s.AddMark(3)
	// s.AddMark(5)
	// s.AddMark(4)
	// a := s.Average()
	// fmt.Println(s.GetName())
	// fmt.Println(s.marks)
	// fmt.Println(a)
	// fmt.Printf("%+v\n", s)
	studs := generateStudents(10000)
	studs2 := make([]*Student, 10000)
	copy(studs2, studs)

	startTime := time.Now()
	bubbleSort(studs)
	result := time.Since(startTime)
	fmt.Printf("bubble: %s\n", result)

	startTime = time.Now()
	sorter := &studentSorter{studs2}
	sort.Sort(sorter)

	result = time.Since(startTime)
	fmt.Printf("quick: %s\n", result)

}
