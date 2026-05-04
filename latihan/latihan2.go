package latihan

import (
	"fmt"
	"sync"
)

//Soal 2 — Channel
//Buat 2 goroutine:
//
//Goroutine pengirim: kirim angka 1–5 ke channel satu per satu
//Goroutine penerima: terima dari channel, lalu print "Menerima angka: X"
//
//Main harus tunggu keduanya selesai.

func Pengirim(channel chan<- int, group *sync.WaitGroup) {
	defer close(channel)
	defer group.Done()
	for i := 1; i <= 5; i++ {
		channel <- i
		fmt.Printf("Mengirim angka: %d\n", i)
	}

}

func Penerima(channel <-chan int, group *sync.WaitGroup) {
	defer group.Done()
	for data := range channel {
		fmt.Printf("Menerima angka: %d\n", data)
	}
}

func Latihan2() {
	channel := make(chan int)
	group := &sync.WaitGroup{}
	group.Add(2)
	go Penerima(channel, group)
	go Pengirim(channel, group)
	group.Wait()
}
