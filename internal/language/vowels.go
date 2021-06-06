package language

import "unicode"

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

func IsVowel(symbol rune) bool {
	for i := 0; i < len(vowels); i++ {
		if unicode.ToUpper(vowels[i]) == unicode.ToUpper(symbol) {
			return true
		}
	}

	return false
}