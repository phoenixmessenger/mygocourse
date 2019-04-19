package dconv

func turnMass (A string) string { 
	var s = []rune(A)
	var f, b int
	f = 0
	b = len(s) - 1
	for f < b {
		s[f], s[b] = s[b], s[f]
		f++
		b--
	}
	return string (s)
}