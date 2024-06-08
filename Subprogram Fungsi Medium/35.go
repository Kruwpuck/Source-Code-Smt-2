// Buatlah fungsi saja
func jumlah(a, b int) int{
	var i, jumlah int
	for i = a; i <= b; i++{
	jumlah = jumlah + i
	}
	if a > b{
	for i = a; i >= b; i--{
	jumlah = jumlah + i
   
	}
	}
	return jumlah
   }
   