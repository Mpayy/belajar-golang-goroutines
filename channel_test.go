package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello Channel"
		fmt.Println("Selesi mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeParameter(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello Channel Parameter"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeParameter(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello Channel In & Out"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
	fmt.Println("Done Channel InOut")
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	go func() {
		channel <- "Hello Channel 1"
		channel <- "Hello Channel 2"
		channel <- "Hello Channel 3"
		fmt.Println(len(channel))
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Done Channel Buffered")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke" + strconv.Itoa(i+1)
		}
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Done Channel Range")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeParameter(channel1)
	go GiveMeParameter(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeParameter(channel1)
	go GiveMeParameter(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
