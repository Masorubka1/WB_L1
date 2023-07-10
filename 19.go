package main

import (
	"fmt"
)

func reverseString(str string) string {
	runes := []rune(str)
	length := len(runes)

	for i := 0; i < length/2; i += 1 {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	return string(runes)
}

func main19() {
	str := "главрыба — абырвалг"
	reversedStr := reverseString(str)

	fmt.Println("Original string:", str)
	fmt.Println("Reversed string:", reversedStr)
}
