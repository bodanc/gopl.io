package word2

import "testing"

func TestIsPalindrome(t *testing.T) {

	// "table-driven testing style"
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"abc", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
	}

	// assertion logic
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}

}
