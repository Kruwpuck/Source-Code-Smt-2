// implementasikan fungsinya saja
func pertemuan(x, y int ) int {
	/* mengembalikan jumlah pertemuan sesuai ketentuan soal */
	 var i, jumlah int
	 for i = 1; i <= 365; i++{
	 if i%x == 0 || i%y == 0{
	 jumlah = jumlah + 1
	 }
	 }
	 return jumlah
	}