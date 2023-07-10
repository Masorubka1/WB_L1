package main

import "fmt"

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main23() {
	slice := []int{1, 2, 3, 4, 5}
	index := 2

	fmt.Println("Original slice:", slice)

	// Удаление i-го элемента
	slice = removeElement(slice, index)

	fmt.Println("Modified slice:", slice)
}
