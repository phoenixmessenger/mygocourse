package main

type studentSorter struct {
	students []*Student
}

func (s *studentSorter) Less(i, j int) bool {
	return s.students[i].Average() <
		s.students[j].Average()
}
func (s *studentSorter) Len() int {
	return len(s.students)
}

func (s *studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]

}
func bubbleSort(studs []*Student) {
	for i := 0; i < len(studs)-1; i++ {
		for j := i + 1; j < len(studs); j++ {
			iavg := studs[i].Average()
			javg := studs[j].Average()
			if iavg < javg {
				tmp := studs[i]
				studs[i] = studs[j]
				studs[j] = tmp
			}
		}
	}
}
