package main

import (
	"fmt"
	"math/big"
)

func main22() {
	// Создаем большие числовые переменные
	a := big.NewInt(1)
	b := big.NewInt(1)

	// Устанавливаем значения a и b как 2^20 + 1
	a.Exp(big.NewInt(2), big.NewInt(20), nil)
	b.Add(a, big.NewInt(1))

	// Умножение
	mul := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a.String(), b.String(), mul.String())

	// Деление
	div := new(big.Int).Div(mul, a)
	fmt.Printf("Деление: %s / %s = %s\n", mul.String(), a.String(), div.String())

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s + %s = %s\n", a.String(), b.String(), sum.String())

	// Вычитание
	sub := new(big.Int).Sub(sum, b)
	fmt.Printf("Вычитание: %s - %s = %s\n", sum.String(), b.String(), sub.String())
}
