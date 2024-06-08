// implementasikan fungsinya saja
func onlineFoodOrdering(f, d int, t bool) int {
	/* mengembalikan total uang yang harus dibayar */
	 var hasil int
	 if t == true{
	 hasil = f + d + 5000
	 }else {
	 hasil = f + d
	 }
	 return hasil
	}