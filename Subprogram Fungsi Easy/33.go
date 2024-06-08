// Buatlah fungsi saja

func jumKendaraan(jumPenumpang, kapKendaraan int) int {
	/* fungsi mengembalikan jumlah kendaraan yang diperlukan
	   jika diinputkan jumlah penumpang dan kapasitas kendaraan */
	if jumPenumpang%kapKendaraan != 0 {
		return jumPenumpang/kapKendaraan + 1
	} else {
		return jumPenumpang / kapKendaraan
	}

}