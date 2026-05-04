package latihan

import (
	"fmt"
	"sync"
)

//Soal 3 — Goroutine + Channel untuk hasil kerja
//Buat function hitungKuadrat(angka int, hasil chan int) yang menghitung kuadrat sebuah angka lalu kirim hasilnya ke channel.
//Jalankan 5 goroutine sekaligus untuk angka 1–5, kumpulkan semua hasilnya, lalu print totalnya.

//Kuadrat 1 = 1
//Kuadrat 2 = 4
//...
//Total = 55

func HitungKuadrat(angka int, hasil chan int, group *sync.WaitGroup) {
	defer group.Done()
	hasil <- angka * angka
}

func Latihan3() {
	hasil := make(chan int)
	group := &sync.WaitGroup{}

	group.Add(5)
	for i := 1; i <= 5; i++ {
		go HitungKuadrat(i, hasil, group)
	}

	go func() {
		group.Wait()
		close(hasil)
	}()

	total := 0
	for result := range hasil {
		fmt.Println(result)
		total += result
	}

	fmt.Println("Total =", total)
}
