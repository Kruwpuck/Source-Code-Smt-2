func findMinMax(bilangan, max, min int) {
	/*   IS : Terdefinisi bilangan bulat positif yang diakhiri bilangan 0 sebagai sentinel.
	FS : Menampilkan bilangan terbesar dan terkecil*/
	fmt.Scan(&bilangan)
	// deklarasi kamus
	max = 0
	min = bilangan
	for bilangan != 0 { // gunakan pengulangan dengan cek kondisi menggunakan operator perbandingan
		// pengecekan bilangan max dan min
		if bilangan > max {
			max = bilangan
		}
		if bilangan <= min {
			min = bilangan
		}
		fmt.Scan(&bilangan)
	}
	fmt.Println(max, min)
}