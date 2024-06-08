// Buatlah fungsi saja

func Jumlah(n, m int) int {
	/*
		Fungsi mengembalikan nilai berupa jumlah bilangan bulat dengan input berupa
		bilangan awal dan bilangan akhir.
	*/
	var x int
	for i := n; i <= m; i++ {
		x += i
	}
	return x
}