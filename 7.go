package main

import (
	"fmt"
	"sync"
)

type SafeMap1 struct {
	sync.RWMutex
	Map map[string]int
}

func (sm *SafeMap1) Add(key string, value int) {
	sm.Lock()
	defer sm.Unlock()

	sm.Map[key] = value
}

func (sm *SafeMap1) Get(key string) (int, bool) {
	sm.RLock()
	defer sm.RUnlock()

	value, ok := sm.Map[key]
	return value, ok
}

func main7_1() {
	safeMap := SafeMap1{
		Map: make(map[string]int),
	}

	var wg sync.WaitGroup

	// Количество горутин для конкурентной записи данных
	numWorkers := 10
	wg.Add(numWorkers)

	// Запуск горутин для записи данных
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()

			// Генерация уникального ключа и значения для записи
			key := fmt.Sprintf("key-%d", workerID)
			value := workerID

			// Запись данных в безопасный map
			safeMap.Add(key, value)
		}(i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод данных из map
	for key, value := range safeMap.Map {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

type SafeMap2 struct {
	Map     map[string]int
	readCh  chan mapRequest
	writeCh chan mapRequest
}

type mapRequest struct {
	operation string
	key       string
	value     int
	resultCh  chan int
}

func NewSafeMap() *SafeMap2 {
	sm := &SafeMap2{
		Map:     make(map[string]int),
		readCh:  make(chan mapRequest),
		writeCh: make(chan mapRequest),
	}
	go sm.processRequests()
	return sm
}

func (sm *SafeMap2) Add(key string, value int) {
	resultCh := make(chan int)
	sm.writeCh <- mapRequest{operation: "add", key: key, value: value, resultCh: resultCh}
	<-resultCh
}

func (sm *SafeMap2) Get(key string) (int, bool) {
	resultCh := make(chan int)
	sm.readCh <- mapRequest{operation: "get", key: key, resultCh: resultCh}
	value := <-resultCh
	_, ok := sm.Map[key]
	return value, ok
}

func (sm *SafeMap2) processRequests() {
	for {
		select {
		case req := <-sm.readCh:
			value, ok := sm.Map[req.key]
			if ok {
				req.resultCh <- value
			} else {
				req.resultCh <- 0
			}

		case req := <-sm.writeCh:
			switch req.operation {
			case "add":
				sm.Map[req.key] = req.value
				req.resultCh <- 0

			// Handle other write operations here

			default:
				req.resultCh <- 0
			}
		}
	}
}

func main7_2() {
	safeMap := NewSafeMap()

	var wg sync.WaitGroup

	// Количество горутин для конкурентной записи данных
	numWorkers := 10
	wg.Add(numWorkers)

	// Запуск горутин для записи данных
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()

			// Генерация уникального ключа и значения для записи
			key := fmt.Sprintf("key-%d", workerID)
			value := workerID

			// Запись данных в безопасный map
			safeMap.Add(key, value)
		}(i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод данных из map
	for key, value := range safeMap.Map {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

func main7() {
	fmt.Println("1:")
	main7_1()
	fmt.Println("\n2:")
	main7_2()
}
