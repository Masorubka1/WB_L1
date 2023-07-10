package main

import "fmt"

func generateNumbers(numbers []int, outCh chan<- int) {
	defer close(outCh)
	for _, num := range numbers {
		outCh <- num
	}
}

func doubleNumbers(inCh <-chan int, outCh chan<- int) {
	defer close(outCh)
	for num := range inCh {
		outCh <- num * 2
	}
}

func printNumbers(inCh <-chan int) {
	for num := range inCh {
		fmt.Println(num)
	}
}

func main9() {
	numbers := []int{1, 2, 3, 4, 5}

	numCh := make(chan int)
	doubleCh := make(chan int)

	go generateNumbers(numbers, numCh)
	go doubleNumbers(numCh, doubleCh)
	printNumbers(doubleCh)
}
