package golang_goroutine

import (
	"fmt"
	"strconv"
	"sync"
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


// 8. Select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {	
		select {
		case data := <- channel1:
			fmt.Println("Data Dari channel 1 ", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data Dari channel 2 ", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// 9. Default Select
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {	
		select {
		case data := <- channel1:
			fmt.Println("Data Dari channel 1 ", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data Dari channel 2 ", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}
		if counter == 2 {
			break
		}
	}
}

// 10. Race Condition
func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Nilai X : ", x)
}

// 11. sync.Mutex
func TestMutexRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	
	time.Sleep(2 * time.Second)
	fmt.Println("Nilai X : ", x)
}

// 12. RWMutex
type BankAccount struct {
	RWMutex sync.RWMutex
	Balace int
}

func (account *BankAccount) AddBelance(amount int) {
	account.RWMutex.Lock()
	account.Balace = account.Balace + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBelance() int {
	account.RWMutex.RLock()
	balance := account.Balace
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {

	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBelance(1)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Balace : ", account.GetBelance())

}

// 14. WaitGroup
func RunAsynchronous(group *sync.WaitGroup, i int) {
	defer group.Done()

	group.Add(1) 

	fmt.Println("Halo Said ", i)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go RunAsynchronous(group , i)	
	}

	group.Wait()
}

// 15. Once
var counter = 1

// tidak boleh ada parameter
func OnlyOnce(){
	counter++
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce) // hanya akan diesekusi sekali
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter : ", counter)
}




