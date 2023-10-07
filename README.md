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

