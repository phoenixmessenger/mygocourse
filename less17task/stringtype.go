package main

type StringArray struct {
	Strings []string
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
