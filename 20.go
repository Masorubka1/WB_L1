package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	words := strings.Fields(str)
	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	return strings.Join(reversedWords, " ")
}

func main20() {
	str := "snow dog sun — sun1 dog1 snow1"
	reversedStr := reverseWords(str)

	fmt.Println("Original string:", str)
	fmt.Println("Reversed string:", reversedStr)
}
