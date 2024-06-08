// Buatlah fungsi saja

func volumeBola(radius float64) float64 {
	/* fungsi menerima input berupa radius dengan tipe data real
	   dan mengembalikan nilai volume bola dengan tipe data real */
	const pi float64 = 3.14
	var hasil float64
	hasil = 4.0 / 3.0 * pi * radius * radius * radius
	return hasil
}
