package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func main18() {
	counter := Counter{value: 0}
	var wg sync.WaitGroup

	// Запускаем горутины для инкрементации счетчика
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Final counter value:", counter.value)
}
