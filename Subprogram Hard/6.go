func totalTabungan(x, y, z int) int {
	/*
		I.S. terdefinisi tiga buah bilangan bulat yang menyatakan uang yang ditabung pertama kali, uang tetap, dan jumlah hari.
		F.S. total tabungan yang terkumpul selama z hari.
	*/

	totalTabungan := 0

	for i := 1; i <= z; i++ {
		if i%2 == 0 || i%3 == 0 {
			totalTabungan += x
			x += y
		}
	}

	return totalTabungan
}
