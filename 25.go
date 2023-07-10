package main

import (
	"fmt"
	"sync"
	"time"
)

type Sleeper struct {
	cond *sync.Cond
}

func NewSleeper() *Sleeper {
	return &Sleeper{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (s *Sleeper) Sleep(duration time.Duration) {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	timer := time.NewTimer(duration)
	defer timer.Stop()

	s.cond.Wait()
}

func (s *Sleeper) WakeUp() {
	s.cond.Signal()
}

func main25() {
	sleeper := NewSleeper()

	go func() {
		fmt.Println("Start")
		time.Sleep(2 * time.Second)
		sleeper.WakeUp()
	}()

	fmt.Println("Sleeping")
	sleeper.Sleep(time.Second)
	fmt.Println("Awake")
}
