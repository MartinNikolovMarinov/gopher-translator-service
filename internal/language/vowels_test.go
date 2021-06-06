package language

import "testing"

func TestIsVowel(t *testing.T) {
	cases := []struct {
		in   rune
		want bool
	}{
		{in: ' ', want: false},
		{in: '?', want: false},
		{in: 'a', want: true},
		{in: 'E', want: true},
		{in: 'i', want: true},
		{in: 'O', want: true},
		{in: 'u', want: true},
		{in: 'j', want: false},
		{in: 'Z', want: false},
	}

	for _, tc := range cases {
		actual := IsVowel(tc.in)
		if actual != tc.want {
			t.Fatalf("IsVowel test failed for %c expected %t got %t", tc.in, tc.want, actual)
		}
	}
}
