package main

import (
	"fmt"

	"github.com/beltranbot/go-testing/services"
)

func main() {
	fmt.Println(services.StringsService.IsPalindrome("anitalavalatina"))
}
