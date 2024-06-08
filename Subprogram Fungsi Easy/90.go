// Buatlah fungsi saja

func jumlahKendaraan(penumpang int) int {
	// Mengembalikan jumlah kendaraan yang harus disewa jika diberikan
	// jumlah penumpang
	const KAPASITAS int = 7
	var kendaraan int
	kendaraan = penumpang / KAPASITAS
	if penumpang%KAPASITAS != 0 {
		kendaraan += 1
	}
	return kendaraan
}