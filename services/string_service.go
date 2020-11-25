package services

import (
	"github.com/beltranbot/go-testing/stringutils"
)

var (
	// StringsService service
	StringsService stringsServiceInterface
)

func init() {
	StringsService = &stringService{}
}

type stringService struct {
	Name string
}

type stringsServiceInterface interface {
	IsPalindrome(someText string) bool
}

// IsPalindrome func
func (s *stringService) IsPalindrome(text string) bool {
	return stringutils.IsPalindrome(text)
}
