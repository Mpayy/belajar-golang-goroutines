package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	pool.Put("Rifaih")
	pool.Put("Aminu")
	pool.Put("Mita")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			defer pool.Put(data)
			time.Sleep(1 * time.Second)
			fmt.Println(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
