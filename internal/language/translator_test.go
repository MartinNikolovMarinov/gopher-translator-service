package language

import "testing"

func TestTranslateWord(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "apple", want: "gapple"},
		{in: "xray", want: "gexray"},
		{in: "chair", want: "airchogo"},
		{in: "square", want: "aresquogo"},
		{in: "nmkty", want: "nmktyogo"},
		{in: "", want: ""},
		{in: "nmpl", want: "nmplogo"}, // words with no vowels remain the same with ogo at the end ?
	}

	for _, tc := range cases {
		actual := TranslateWord(tc.in)
		if actual != tc.want {
			t.Fatalf("Word Translate test failed for %s expected %s got %s", tc.in, tc.want, actual)
		}
	}
}

func TestTranslateSentance(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"apple", "xray", "chair", "square", "nmkty"}, "gapple gexray airchogo aresquogo nmktyogo"},
		{[]string{"nmpl"}, "nmplogo"},
		{[]string{}, ""},
		{nil, ""},
	}

	for _, tc := range cases {
		actual := TranslateSentence(tc.in)
		if actual != tc.want {
			t.Fatalf("Word Translate test failed for %s expected %s got %s", tc.in, tc.want, actual)
		}
	}
}
