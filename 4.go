package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func handleSignals4(dataCh *chan int, wg *sync.WaitGroup) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalCh
		// При получении сигнала завершения программы
		// закрываем канал и ожидаем завершения всех воркеров
		close(*dataCh)
		wg.Wait()
		os.Exit(0)
	}()
}

func worker4(id int, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataCh {
		fmt.Printf("Worker %d received: %d\n", id, data)
		// Выполнение задачи воркера
		// ...
	}
}

func main4() {
	// Определение флага командной строки для количества воркеров
	// go run main.go -workers=10
	numWorkers := flag.Int("workers", 5, "number of workers")
	flag.Parse()

	// Создаем канал для передачи данных
	dataCh := make(chan int)

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Запускаем воркеры
	for i := 0; i < *numWorkers; i++ {
		wg.Add(1)
		go worker4(i, dataCh, &wg)
	}

	// Обработка сигналов завершения программы
	handleSignals4(&dataCh, &wg)

	// Постоянная запись данных в канал
	for i := 1; ; i++ {
		// recover - долго, потому что создаётся стектрейс
		// case dataCh <- i == case _, ok := dataCh <- i; ok
		select {
		case dataCh <- i:
			// Отправка данных в канал
		case <-dataCh:
			// Канал закрыт, выход из цикла
			return
		}
	}

}
