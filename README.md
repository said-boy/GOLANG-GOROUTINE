1. goroutine
    membuat function dijalankan secara asynchronous, caranya dengan menambahkan keyword `go <function>`. contoh : go HelloWorld()
    
    Catatan: goroutine tidak cocok jika digunakan di function yang ada return value, karena goroutine tidak dapat menangkap return valuenya. 

2. goroutine sangat ringan, dapat membuat banyak tanpa harus pusing  pemakaian memori.

3. channel
    channel adalah tempat menyimpan data dari goroutine yang dapat diambil. 