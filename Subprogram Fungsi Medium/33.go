// implementasikan fungsinya saja
func hitungDigit(x int ) int {
	/* mengembalikan jumlah digit dari bilangan masukan */
	 var jumlah int
	 for x != 0{
	 x = x / 10
	 jumlah = jumlah + 1
	 }
	 return jumlah
	}