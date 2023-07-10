package main

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

func myGoroutine6_1(wg *sync.WaitGroup) {
	defer wg.Done()

	// Выполнять действия в горутине

	// Завершить выполнение горутины
	return
}

func main6_1() {
	var wg sync.WaitGroup

	wg.Add(1)
	go myGoroutine6_1(&wg)

	// Ожидать завершения горутины
	wg.Wait()
}

func myGoroutine6_2(stopCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Выполнять действия в горутине

	// Проверять канал на наличие сигнала остановки
	select {
	case <-stopCh:
		// Получен сигнал остановки, завершить выполнение горутины
		return
	default:
		// Продолжить выполнение действий
	}
}

func main6_2() {
	stopCh := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go myGoroutine6_2(stopCh, &wg)

	// Через некоторое время отправить сигнал остановки
	go func() {
		time.Sleep(time.Second)
		close(stopCh)
	}()

	// Ожидать завершения горутины
	wg.Wait()
}

func myGoroutine6_3(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// Выполнять действия в горутине

	// Проверять контекст на наличие сигнала остановки
	select {
	case <-ctx.Done():
		// Получен сигнал остановки, завершить выполнение горутины
		return
	default:
		// Продолжить выполнение действий
	}
}

func main6_3() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go myGoroutine6_3(ctx, &wg)

	// Через некоторое время вызвать cancel для остановки горутины
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	// Ожидать завершения горутины
	wg.Wait()
}

var stopFlag int32

func myGoroutine6_4(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Проверять значение флага остановки
		if atomic.LoadInt32(&stopFlag) == 1 {
			// Флаг остановки установлен, завершить выполнение горутины
			return
		}

		// Выполнять действия в горутине
	}
}

func main6_4() {
	var wg sync.WaitGroup

	wg.Add(1)
	go myGoroutine6_4(&wg)

	// Через некоторое время установить флаг остановки
	go func() {
		time.Sleep(time.Second)
		atomic.StoreInt32(&stopFlag, 1)
	}()

	// Ожидать завершения программы
	wg.Wait()
}

var mutex sync.Mutex
var stopSignal bool

func myGoroutine6_5(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Захватить блокировку перед проверкой флага остановки
		mutex.Lock()
		if stopSignal {
			// Флаг остановки установлен, освободить блокировку и завершить выполнение горутины
			mutex.Unlock()
			return
		}
		mutex.Unlock()

		// Выполнять действия в горутине
	}
}

func main6_5() {
	var wg sync.WaitGroup

	wg.Add(1)
	go myGoroutine6_5(&wg)

	// Через некоторое время установить флаг остановки
	go func() {
		time.Sleep(time.Second)
		mutex.Lock()
		stopSignal = true
		mutex.Unlock()
	}()

	// Ожидать завершения программы
	wg.Wait()
}

func main6() {
	main6_1()
	main6_2()
	main6_3()
	main6_4()
	main6_5()
}
