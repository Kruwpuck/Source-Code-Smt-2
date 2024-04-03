// Buatlah fungsi saja

func jumlahDigit(n int) int {
	// fungsi akan mengembalikan jumlah digit-digit bilangan
	var d1, d2 int
	d1 = n / 10
	d2 = n % 10
	return d1 + d2
}