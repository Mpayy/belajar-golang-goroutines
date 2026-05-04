package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			defer group.Done()
			once.Do(OnlyOnce)
		}()
	}
	group.Wait()
	fmt.Println("Counter: ", counter)
}
