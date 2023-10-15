package golang_goroutine

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
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

// 16. Pool
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "Default" // default value pengganti <nil>
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Muhammad Said") // Meletakkan data pada pool
	pool.Put("Alkhudri")

	for i := 0; i < 10; i++ {
		go func(j int){
			group.Add(1)
			data := pool.Get() // mengambil data dari pool
			fmt.Println("Goroutine ke : ",j," data : ",data)
			time.Sleep(1 * time.Second)
			pool.Put(data) // meletakkan kembali data pada pool
			group.Done()
		}(i)
	}
	group.Wait()

}

// 17. Map

func TestMap(t *testing.T) {

	data := sync.Map{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func(number int) {
			group.Add(1)

			data.Store(number, number) // untuk memasukkan data ke map

			group.Done()
		}(i)
	}

	group.Wait()

	// mengambil data dari map dengan iteration
	data.Range(func(key, value any) bool {
		fmt.Println(key , " : ", value)
		return true // true = untuk terus mengambil data selanjutnya , false = untuk data pertama saja.
	})

}

// 18. Cond
var Mutex = sync.Mutex{}
var Cond = sync.NewCond(&Mutex) // harus mengisi lockernya. 
var Group = sync.WaitGroup{}

func WaitCondition(number int){
	defer Group.Done()

	Cond.L.Lock() // mengunci goroutine.
	
	Cond.Wait() // menunggu sinyal diberikan
	/* 
		jika tidak ada signal yang diberikan akan terjadi error
		karena goroutine nya menunggu terus. (deadlock!)
	*/
	fmt.Println("Goroutine ke : ", number)

	Cond.L.Unlock()

}

func TestCond(t *testing.T) {

	// menjalankan sebagai goroutine
	for i := 0; i < 10; i++ {
		Group.Add(1)
		go WaitCondition(i)
	}
	// Signal / broadcast harus dijalankan secara goroutine
	// dengan jumlah yang sama seperti go Cond nya.

	// mengirim signal sebagai goroutine ke Cond
	// keluar satu persatu
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1*time.Second)
			Cond.Signal()	
		}
	}()

	// keluar bersama sama
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	Cond.Broadcast()
	// 	}()
		
	Group.Wait()
	fmt.Println("Selesai..")

}

// 19. Atomic
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {

		group.Add(1) // ini yang benar (sebelum goroutine dijalankan)

		go func() {
			atomic.AddInt64(&x, 1)
			group.Done() // diberikan didalam goroutinenya
		}()

	}

	group.Wait()
	fmt.Println(x)

}

// 20. Timer
func TestTimer(t *testing.T) {
	// Manual
	Manualtimer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())
	timeManual := <- Manualtimer.C
	fmt.Println(timeManual)
	
	// Auto = otomatis mengambil channelnya saja tanpa .C
	AutoTimer := time.After(5 * time.Second)
	timeAuto := <- AutoTimer
	fmt.Println(timeAuto)

	// otomatis menjalankan function ketika time terpenuhi
	Group.Add(1)
	// AfterFunc berjalan secara async
	time.AfterFunc(5 * time.Second, func(){
		fmt.Println(time.Now()) // akan dijalankan setelah 5 detik
		Group.Done()
	})
	Group.Wait()

}

// 21. Ticker
func TestTickerTick(t *testing.T) {
	fmt.Println("Ticker dimulai...")
	ticker := time.NewTicker(1 * time.Second)
	// ticker akan terus berulang, jadi saya membatasinya sebanyak 5 saja.
	for i := 0; i < 5; i++ {
		dataTicker := <- ticker.C
		fmt.Println(dataTicker)
	}
	ticker.Stop() // untuk menghentikan ticker
	fmt.Println("Ticker dihentikan...")
	
	fmt.Println("")

	// Tick sama seperti Ticker hanya saja
	// dia langsung mengembalikan channelnya, tanpa mengakses .C
	fmt.Println("Tick diMulai...")
	tick := time.Tick(1 * time.Second)
	for i := 0; i < 5; i++ {
		dataTick := <- tick // Tick langsung tanpa .C (channel)
		fmt.Println(dataTick)
	}
	fmt.Println("Tick dihentikan...")
}

func TestGomaxProcs(t *testing.T) {
	
	for i := 0; i < 10; i++ {
		Group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			Group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total Cpu : ",totalCpu)
	
	// runtime.GOMAXPROCS(20) // untuk mengubah thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread : ",totalThread)
	
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine : ",totalGoroutine)

	Group.Wait()
}

