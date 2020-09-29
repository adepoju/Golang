package word

import "strings"

// UseCount returns the number of times a word appears in a string
func UseCount( s string) map[string]int {
	sx := strings.Fields(s)

	m := make(map[string]int)
	for _,v := range sx {
		m[v]++ 
	}
	return m
}

// Count returns the number of words in a string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(s)
}
