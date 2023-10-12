1. goroutine
    membuat function dijalankan secara asynchronous, caranya dengan menambahkan keyword `go <function>`. contoh : go HelloWorld()
    
    Catatan: goroutine tidak cocok jika digunakan di function yang ada return value, karena goroutine tidak dapat menangkap return valuenya. 

2. goroutine sangat ringan, dapat membuat banyak tanpa harus pusing  pemakaian memori.

3. channel
    channel adalah tempat menyimpan data dari goroutine yang dapat diambil. 

6. Buffered Channel
    - adalah menambahkan ruang pada channel, kita bisa menambahkan ruang lebih dari 1, defaultnya 0. 
    - 0 itu berarti data yang bisa diisi cuma 1 data dan data tersebut juga harus digunakan (diambil). jika tidak maka akan error deadlock!.
    - dengan buffer memungkinkan kita mengisi data lebih dari satu dan tidak akan terjadi kesalahan apabila data tersebut tidak diambil

7. Range Channel
    - Daripada kita mengakses satu persatu data di channel yang belum tentu ada berapa banyak data, lebih baik menggunakan perulangan saja.
    - tetapi ada syarat dimana setelah kita looping memasukkan data channel harus di close, agar bisa menjalankan iterasi atau perulangan pada pengambilan data.
    - jika anda tidak menambahkan buffer pada channel nya, maka perulanganya harus dijalankan secara asyn dan harus di close channel diakhir perulangan.

8. Select channel
    - dapat mengambil data antara channel yang tercepat.
    - jumlah case pada select harus sesuai dengan banyaknya channel yang anda punya.

9. Default Select
    - Akan menunggu data dari sebuah channel sambil melakukan sesuatu hingga data tersebut ada pada channel, jika tidak ada data atau belum ada data maka akan terus ditunggu.

10. Race Condition
    - race condition adalah problem dimana ketika kita mengsharing sebuah variabel terhadap goroutine.

11. sync.Mutex
    - untuk mengatasi race condition bisa menggunakan ini. dimana goroutine yang sejatinya berjalan secara asyn akan dibuat mengantri dalam mengakses variabel yang disharing.
    - sync.Mutex ini akan membuat variabel yang disharing terkunci untuk 1 goroutine yang tercepat dan setelah goroutine itu selesai maka variabel nya akan terbuka lagi untuk umum.

12. RWMutex
    - Read dan Write Mutex

13. Deadlock
    - adalah error dimana proses berjalan terus menerus. dalam goroutine contoh kasusnya adalah ketika satu goroutine menunggu goroutine lainnya.
    - error ini juga bisa terjadi saat menggunakan mutex. jadi harus hati" saat menggunakan mutex.

14. WaitGroup
    - gunakan wait group untuk menunggu semua goroutine selesai.

15. Once
    - digunakan untuk memastikan sebuah function di esekusi hanya sekali, selebihnya akan dihiraukan.
    - hanya goroutine pertama yang dapat mengakses function tersebut.

16. Pool
    - seperti tempat untuk menyimpan data mirip seperti array atau slice. pool ini digunakan untuk menyimpan data yang mahal akan diambil oleh goroutine.
    - pool ini mengelola memanfaatkan ulang objek. untuk pengurangan penggunaan memori.




