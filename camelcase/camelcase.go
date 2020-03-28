package camelcase

import "unicode"

func CountWords(word string) int {
	count := 1
	for _, r := range word {
		if unicode.IsUpper(r) {
			count++
		}
	}
	return count
}
