// Buatlah fungsi saja

func hargaBarangKenaPajak(negri string, harga int) float64 {
	// Mengembalikan harga barang setelah terkena pajak
	// Barang dalam negeri terkena pajak 8%, luar negeri 18%
	var total float64
	if negri == "dalam" {
		total = float64(harga) - (float64(harga) * 8 / 100)
	} else if negri == "luar" {
		total = float64(harga) - (float64(harga) * 18 / 100)
	}
	return total
}