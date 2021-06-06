package validators

import "unicode"

func IsValidEnglishWord(word string) bool {
	for _, r := range word {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func IsValidEnglishSentence(words []string) bool {
	for i := 0; i < len(words); i++ {
		if !IsValidEnglishWord(words[i]) {
			return false
		}
	}

	return true
}