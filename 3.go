package main

import (
	"fmt"
	"runtime"
	"sync"
)

func calcSquare3(start, finish int, array []int, wg *sync.WaitGroup) {
	defer wg.Done()
	array[start] *= array[start]
	for i := start + 1; i < finish; i += 1 {
		array[start] += array[i] * array[i]
	}
}

func main3() {
	numCPU := runtime.NumCPU()
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var wg sync.WaitGroup

	chunkSize := len(array) / numCPU

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		start := i * chunkSize
		finish := (i + 1) * chunkSize

		if i == numCPU-1 {
			finish = len(array)
		}

		go calcSquare3(start, finish, array, &wg)
	}

	wg.Wait()

	ans := 0
	if chunkSize == 0 {
		ans = array[0]
	} else {
		for i := 0; i < numCPU && chunkSize != 0; i += 1 {
			ans += array[i*chunkSize]
		}
	}
	fmt.Println(array)
	fmt.Println(ans)
}
