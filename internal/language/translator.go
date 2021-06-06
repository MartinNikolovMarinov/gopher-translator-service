package language

import "strings"

func TranslateWord(word string) string {
	if word == "" {
		return ""
	}

	wordRunes := []rune(word)
	if res := applyPreTranslationRules(wordRunes); res != "" {
		return res
	}

	for i := 0; i < len(wordRunes); i++ {
		firstChar := wordRunes[0]
		nextChar := wordRunes[1]
		if IsVowel(firstChar) {
			break
		}

		consonantFollowedByQU := (i < len(wordRunes) - 1) && (firstChar == 'q') && (nextChar == 'u')
		if consonantFollowedByQU {
			// special case we need to rotate twice:
			rotateLeftByOne(wordRunes)
		}

		rotateLeftByOne(wordRunes)
	}

	return string(wordRunes) + "ogo"
}

func TranslateSentence(words []string) string {
	var sb strings.Builder
	for i := 0; i < len(words); i++ {
		translated := TranslateWord(words[i])
		if i > 0 { // don't write space in front of the final result
			_ = sb.WriteByte(byte(' '))
		}
		_, _ = sb.WriteString(translated)
	}

	return sb.String()
}

func applyPreTranslationRules(wordRunes []rune) string {
	firstLetterIsVowel := IsVowel(wordRunes[0])
	if firstLetterIsVowel {
		return "g" + string(wordRunes)
	}
	startsWithXR := (len(wordRunes) > 1) && (wordRunes[0] == 'x') && (wordRunes[1] == 'r')
	if startsWithXR {
		return "ge" + string(wordRunes)
	}

	return ""
}

func rotateLeftByOne(arr []rune) {
	var i int
	first := arr[0]
	for i = 0; i < len(arr) - 1; i++ {
		arr[i] = arr[i+1]
	}
	arr[i] = first
}