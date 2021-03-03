package word1

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("wow") {
		t.Error(`IsPalindrome("wow") = false`)
	}
	if !IsPalindrome("civic") {
		t.Error(`IsPalindrome("civic") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
	// this should return true, which would get flipped to false (!) and the if statement would be ignored;
/*	if !IsPalindrome("A man, a plan, a canal: Panama") {
		t.Error(`IsPalindrome("A man, a plan, a canal: Panama") = false`)
	}*/
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("bogdan") {
		t.Error(`IsPalindrome("bogdan")`)
	}
}

// this reveals a bug in my program! IsPalindrome should return true in this case (since été is indeed a palindrome);
// the fact that a return value of false gets flipped to true means the code inside the if statement block gets
// executed and 'go test' test will fail as a consequence;
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		// to avoid writing the long input string twice, we use Errorf, which provides formatting like Printf
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
