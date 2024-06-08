func hitungSKS(totalSKS int) int {
	var kodeMatkul string
	fmt.Scan(&kodeMatkul) // Scan Kode mata Kuliah

	if kodeMatkul == "0" {
		return totalSKS
	}

	sks := 0 // Inisialisasi SKS

	if len(kodeMatkul) > 0 { // Pastikan kodeMatkul tidak kosong
		switch kodeMatkul[0] {
		case 'A', 'B', 'C':
			sks = 2
		case 'D', 'E', 'F':
			sks = 3
		}
	}

	return hitungSKS(totalSKS + sks) // Rekursi untuk menghitung total SKS
}