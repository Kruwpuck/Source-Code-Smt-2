// Buatlah fungsi saja

func jumlahDeret(a, b int) int {
	/*
		Fungsi mengembalikan nilai berupa jumlah bilangan bulat dengan input berupa
		bilangan awal "n" dan bilangan akhir "m". Gunakanlah perulangan while-do.
	*/
	var i, jumlah int
	i = a
	for i <= b {
		jumlah += i
		i += 1
	}
	return jumlah
}