package main

import (
	"fmt"
	"strings"
)

func areAllCharactersUnique(str string) bool {
	seen := make(map[rune]bool)

	// Приводим строку к нижнему регистру перед проверкой
	lowerStr := strings.ToLower(str)

	for _, char := range lowerStr {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func main() {
	str1 := "abcd"
	fmt.Printf("Строка '%s' содержит уникальные символы: %t\n", str1, areAllCharactersUnique(str1))

	str2 := "abCdefAaf"
	fmt.Printf("Строка '%s' содержит уникальные символы: %t\n", str2, areAllCharactersUnique(str2))

	str3 := "aabcd"
	fmt.Printf("Строка '%s' содержит уникальные символы: %t\n", str3, areAllCharactersUnique(str3))
}
