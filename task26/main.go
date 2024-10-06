package main

import (
	"fmt"
	"strings"
)

func unique(word string) bool {
	cache := make(map[rune]struct{})
	lowerString := strings.ToLower(word)
	for _, v := range lowerString {

		if _, ok := cache[v]; ok {
			return false
		}
		cache[v] = struct{}{}
	}

	return true
}

func main() {
	letter1 := "abcd"
	fmt.Println(letter1, " - ", unique(letter1))
	letter2 := "abCdefAaf"
	fmt.Println(letter2, " - ", unique(letter2))
}
