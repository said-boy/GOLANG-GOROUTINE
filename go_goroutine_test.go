package golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 1. Gorouttine
func HelloWorld() {
	fmt.Println("Hello World!")
}

func TestHelloWorld(t *testing.T) {
	go HelloWorld() // goroutine (asynchronous)
	fmt.Println("Ups!")
}

// 2. Gouroutine sangat ringan
func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 1; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
} 

// 3. Channel
func TestChannelGoroutine(t *testing.T) {
	// membuat Channel
	channel := make(chan string)
	// menutup channel setelah function selesai dijalankan 
	defer close(channel)

	go func(){
		time.Sleep(2 * time.Second)
		channel <- "Muhammad Said Alkhudri" // mengisi channel dengan string
		fmt.Println("Berhasil Mengirim data ke Channel")
	}()

	data := <- channel // mengambil data dari channel
	fmt.Println(data)
	time.Sleep(3 * time.Second)
}

// 4. channel sebagai parameter
func GiveMeResponse(channel chan string){
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Said Alkhudri"
}

func TestChannelAsParameter(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	time.Sleep(3 * time.Second)

}

// 5. channel send-only and recieve-only
// send-only
func OnlyIn(channel chan<- string) {
	channel <- "Muhammad Said Alkhudri"
}

// receive-only
func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data)
}

func TestChannelInAndOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)

}

// 6. Buffered Channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "Muhammad"
	channel <- "Said"

	fmt.Println(<- channel)
	fmt.Println(<- channel)

}

// 7. Range Channel
func TestRangeChannelWithoutBuffered(t *testing.T) {
	channel := make(chan string) // no buffered

	go func () {
		for i := 0 ; i < 10 ; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel) // harus di close dulu sebelum datanya diambil.
	}()
		
	for data := range channel {
		fmt.Println("Data", data)
	}

}
// 7. Range Channel
func TestRangeChannelWithBuffered(t *testing.T) {
	channel := make(chan string, 10) // buffered

	for i := 0 ; i < 10 ; i++ {
		channel <- "Perulangan ke " + strconv.Itoa(i)
	}
	close(channel) // harus di close dulu sebelum datanya diambil.

	for data := range channel {
		fmt.Println("Data", data)
	}

}
