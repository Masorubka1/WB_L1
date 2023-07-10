package main

import (
	"fmt"
	"time"
)

func sendValues(ch chan<- int, n int) {
	for i := 1; i <= n; i++ {
		ch <- i
		fmt.Println("Sent:", i)
		time.Sleep(time.Second) // Ожидание между отправкой значений
	}

	close(ch) // Закрытие канала после отправки всех значений
}

func receiveValues(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received:", value)
	}
}

func main5() {
	ch := make(chan int)

	N := 5 // Количество значений для отправки

	go sendValues(ch, N) // Запуск горутины для отправки значений в канал
	go receiveValues(ch) // Запуск горутины для чтения значений из канала

	time.Sleep(time.Second * time.Duration(N)) // Ожидание N секунд

	fmt.Println("Program finished")
}
