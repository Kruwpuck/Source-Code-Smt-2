// Buatlah fungsi saja
func mesinKue (terigu, gula, mentega, telur int)int {
	// mengembalikan jumlah kue yang dapat dibuat
	var jumlah int
	 jumlah = terigu / 20
	 if gula / 5 < jumlah {
	 jumlah = gula / 5
	 }
	 if mentega / 6 < jumlah {
	 jumlah = mentega / 6
	 }
	 if telur < jumlah {
	 jumlah = telur
	 }
	 return jumlah
	}
	