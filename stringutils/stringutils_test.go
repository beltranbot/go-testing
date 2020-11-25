package stringutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// executes before all the tests start executing
func TestMain(m *testing.M) {
	// Do whatever you need to run before running the test cases.
	os.Exit(m.Run())
}

// unit test

func TestIsPalindromeLenOne(t *testing.T) {
	assert.True(t, IsPalindrome("a"), "single character should be palindrome")
}

func TestIsPalindromeTwoSameCharacters(t *testing.T) {
	assert.True(t, IsPalindrome("aa"), "the sametwo characters should be palindrome")
}

func TestIsPalindromeNotPalindrome(t *testing.T) {
	assert.False(t, IsPalindrome("abc"), "not palindrome should return false")
}

func TestIsPalindromeNoError(t *testing.T) {
	assert.True(t, IsPalindrome("anitalavalatina"), "palindrome should return true")
}

func TestIsPalindromeWithSpacesNoReplaces(t *testing.T) {
	assert.False(t, isPalindrome("anita lava la tina"), "palindrome without spaces removed should return false")
}

func TestIsPalindromeWithSpacesNoError(t *testing.T) {
	assert.True(t, IsPalindrome("anita lava la tina"), "palindrome should return true")
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("anitalavalatina")
	}
}
