//buatlah fungsi saja

func uangMakan(hadir int) int {
	//mengembalikan jumlah uang makan berdasarkan jumlah kehadiran
	if hadir >= 5 && hadir <= 10 {
		return 220000
	} else if hadir >= 11 && hadir <= 15 {
		return 320000
	} else if hadir >= 16 && hadir <= 20 {
		return 380000
	} else if hadir > 20 {
		return 420000
	} else {
		return 0
	}

}