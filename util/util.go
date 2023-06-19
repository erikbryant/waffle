package util

// Keys returns a slice of keys from the given map.
func Keys(m map[rune]int) []rune {
	p := []rune{}
	for k := range m {
		p = append(p, k)
	}
	return p
}
