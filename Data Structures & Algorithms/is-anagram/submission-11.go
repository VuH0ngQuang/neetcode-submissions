import "slices"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := []byte(s)
	b := []byte(t)

	slices.Sort(a)
	slices.Sort(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		} 
	}

	return true
}
