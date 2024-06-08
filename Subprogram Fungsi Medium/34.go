// implementasikan fungsinya saja
func jumlahGanjil(a, b int) int {
	/* mengembalikan penjumlahan bilangan ganjil dalam interval masukan */
	 var i, jumlah int
	 for i = a; i <= b; i++{
	 if i%2 != 0{
	 jumlah = jumlah + i
	 }
	 }
	 if a > b{
	 for i = a; i >= b; i--{
	 if i%2 != 0{
	 jumlah = jumlah + i
	 }
	 }
	 }
	 return jumlah
	}