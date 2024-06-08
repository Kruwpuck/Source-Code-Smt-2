func isiTanki(tanki, ember, jumlah int) {
	/*
	   IS : terdefinisi kapasitas tanki yang bertipe bilangan bulat, yang akan diisi menggunakan ember secara terus menerus sampai tanki penuh.
	   FS : menampilkan kondisi tanki penuh atau tidak
	*/

	// Input kapasitas tanki
	fmt.Scan(&tanki)

	// Inisialisasi jumlah air di tanki
	jumlah = 0

	// Operasi pengisian tanki
	for {
		fmt.Scan(&ember)
		jumlah = jumlah + ember
		fmt.Println(jumlah >= tanki)
		if jumlah >= tanki {
			break
		}
	}
}
