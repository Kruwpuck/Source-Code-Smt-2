type dataPenjualan [12]int

func inputPenjualan(arrPenjualan *dataPenjualan) {
	/*I.S. Terdefinisi arrPenjualan
	F.S. arrPenjualan berisi dengan data penjualan tiap bulan dari tahun sebelumnya*/
	var jumlah int
	for i := 0; i < 12; i++ {
		fmt.Scan(&jumlah)
		arrPenjualan[i] = jumlah
	}
}

func avgPenjualan(arrPenjualan dataPenjualan) float64 {
	// Mengembalikan rata-rata dari data penjualan
	var total int
	for _, penjualan := range arrPenjualan {
		total += penjualan
	}
	return float64(total) / float64(12)
}

func bulanTerbanyak(arrPenjualan dataPenjualan) int {
	// Mengembalikan bulan dengan total penjualan tertinggi
	var max, bulanMax int
	for bulan, penjualan := range arrPenjualan {
		if penjualan > max {
			max = penjualan
			bulanMax = bulan + 1 // Karena bulan dimulai dari 1, bukan dari 0
		}
	}
	return bulanMax
}

func bulanTersedikit(arrPenjualan dataPenjualan) int {
	// Mengembalikan bulan dengan total penjualan terendah
	min := arrPenjualan[0]
	bulanMin := 1
	for bulan, penjualan := range arrPenjualan {
		if penjualan < min {
			min = penjualan
			bulanMin = bulan + 1 // Karena bulan dimulai dari 1, bukan dari 0
		}
	}
	return bulanMin
}