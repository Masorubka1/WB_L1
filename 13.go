package main

import "fmt"

func swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func main13() {
	x, y := 10, 20
	fmt.Println("Before swap: x =", x, ", y =", y)

	x, y = swap(x, y)
	fmt.Println("After swap: x =", x, ", y =", y)
}
