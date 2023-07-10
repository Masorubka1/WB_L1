package main

import (
	"fmt"
)

func setBit(num *int64, i int, value int) {
	if value == 1 {
		*num |= 1 << i // Устанавливаем i-й бит в 1
	} else {
		*num &= ^(1 << i) // Устанавливаем i-й бит в 0
	}
}

func main8() {
	var num int64 = 0
	fmt.Println("Исходное значение:", num)

	setBit(&num, 3, 1)
	fmt.Println("Установлен 3-й бит в 1:", num)

	setBit(&num, 5, 0)
	fmt.Println("Установлен 5-й бит в 0:", num)
}
