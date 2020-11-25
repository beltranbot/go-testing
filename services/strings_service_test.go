package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	isPalindromeFunction func(t string) bool
)

type serviceMock struct{}

func (s *serviceMock) IsPalindrome(text string) bool {
	return isPalindromeFunction(text)
}

func TestIsPalindrome(t *testing.T) {
	StringsService = &serviceMock{}
	isPalindromeFunction = func(t string) bool {
		return true
	}
	assert.True(t, StringsService.IsPalindrome("ana"))
}
