package main

import (
	"fmt"
	"math/rand"
)

func quicksort(arr []interface{}, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quicksort(arr, low, pivotIndex-1)
		quicksort(arr, pivotIndex+1, high)
	}
}

func partition(arr []interface{}, low, high int) int {
	pivotIndex := rand.Intn(high-low+1) + low
	pivot := arr[pivotIndex]
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	i := low - 1
	for j := low; j < high; j++ {
		if arr[j].(int) < pivot.(int) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main16() {
	arr := []interface{}{9, 5, 1, 8, 3, 2, 7, 6, 4}
	fmt.Println("Original array:", arr)

	quicksort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}
