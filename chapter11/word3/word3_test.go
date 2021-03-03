package word3

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomPalindromes(t *testing.T) {

	// initialize a pseudo-random number generator
	seed := time.Now().UTC().UnixNano()
	t.Logf("random seed: %d", seed)

	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		fmt.Println(rng, p)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}

}
