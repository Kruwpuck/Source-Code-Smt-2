// Buat function saja
func jumlahGanjil(n int) int {
	// mengembalikan jumlah 1 + 3 + 5 + ... N
	var jumlah, i int
	for i = 1; i <= n; i++{
	if i % 2 != 0 {
	jumlah = jumlah + i
	}
	}
	return jumlah
   
   }
   