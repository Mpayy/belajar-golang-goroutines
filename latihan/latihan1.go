package latihan

import (
	"fmt"
	"sync"
)

//Soal 1 — Goroutine Basic
//Buat program yang menjalankan 5 goroutine sekaligus. Setiap goroutine print:

//Worker 1 mulai
//Worker 1 selesai
//Worker 2 mulai
//...

//Pastikan main tidak exit sebelum semua goroutine selesai — pakai sync.WaitGroup.

func Worker(group *sync.WaitGroup, i int) {
	defer group.Done()
	fmt.Printf("Worker %d mulai\n", i)
	fmt.Printf("Worker %d selesai\n", i)
}

func Latihan1() {
	group := &sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		group.Add(1)
		go Worker(group, i)
	}
	group.Wait()
}
