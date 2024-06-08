func parkir (ken string, waktu int) int {
	// mengembalikan total biaya parkir kendaraan dalam durasi tertentu
	var harga int
	if ken == "motor"{
	if waktu == 1{
	harga = 2000
	}else if waktu > 1{
	harga = 2000 + (500*(waktu - 1))
	}
	}else if ken == "mobil"{
	if waktu == 1{
	harga = 3000
	}else if waktu > 1{
	harga = 3000 + (1000*(waktu - 1))
	}
	}
	return harga
}   