// Buatlah fungsi saja

func volumeKerucut(r, t float64) float64 {
	/* fungsi menerima input berupa radius dan tinggi dengan tipe data real
	   dan mengembalikan nilai volume kerucut dengan tipe data real */
	const pi float64 = 3.14
	var hasil float64
	hasil = 1.0 / 3.0 * pi * r * r * t
	return hasil
}