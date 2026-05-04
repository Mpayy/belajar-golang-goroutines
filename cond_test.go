package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var group = sync.WaitGroup{}
var cond = sync.NewCond(&locker)

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	defer cond.L.Unlock()
	cond.Wait()
	fmt.Println("Cond", value)
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// Signal = untuk memberi satu Signal ke Wait untuk jalanin 1 Goroutines
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// Broadcast = untuk memberi satu Signal ke Wait untuk jalanin semua Goroutines
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	cond.Broadcast()
	//}()

	group.Wait()
}
