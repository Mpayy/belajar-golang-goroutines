package latihan

import (
	"fmt"
	"sync"
	"time"
)

//Soal 4 — Real case: proses data paralel
//Kamu punya slice nama:

//names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}

//Simulasikan setiap nama diproses (pakai time.Sleep(100 * time.Millisecond) sebagai simulasi kerja), lalu print "[nama] sudah diproses".
//Jalankan semua secara paralel dan pastikan program tidak exit sebelum semua selesai. Hitung juga total waktu eksekusi — harusnya sekitar 100ms, bukan 500ms.

func ReadName(name string, group *sync.WaitGroup) {
	defer group.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(name, "sudah diproses")
}

func Latihan4() {
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	group := &sync.WaitGroup{}
	start := time.Now()

	//for _, name := range names {
	//	group.Add(1)
	//	go func(name string) {
	//		defer group.Done()
	//		fmt.Println(name, "sudah diproses")
	//		time.Sleep(100 * time.Millisecond)
	//	}(name)
	//}

	for _, name := range names {
		group.Add(1)
		go ReadName(name, group)
	}
	group.Wait()

	finish := time.Now()
	fmt.Println(finish.Sub(start))
}
