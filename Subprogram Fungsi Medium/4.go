// implementasikan fungsinya saja
func sumRepeatUntillTen( n int ) int {
	// mengembalikan hasil penjumlahan dari 1 hingga (10 x N)
	var stop bool
	var jumlah, i int
   
	stop = false
	for !stop{
	if n == 0{
	jumlah = 1
	}
	for i = 0; i <= 10*n;i++{
	jumlah = jumlah + i
	stop = i == 10*n
	}
	}
   
	return jumlah
   }