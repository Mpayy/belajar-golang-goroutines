package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// akan terjadi race condition jika hanya seperti ini
func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x)
}
