package main

import (
	"fmt"
	"runtime"
	"sync"
)

func calcSquare2(start, finish int, array []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i < finish; i += 1 {
		array[i] = array[i] * array[i]
	}
}

func main2() {
	numCPU := runtime.NumCPU()
	array := []int{1, 2, 3, 4, 5, 6, 7}
	var wg sync.WaitGroup

	chunkSize := len(array) / numCPU

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		start := i * chunkSize
		finish := (i + 1) * chunkSize

		if i == numCPU-1 {
			finish = len(array)
		}

		go calcSquare2(start, finish, array, &wg)
	}

	wg.Wait()

	fmt.Println(array)
}
