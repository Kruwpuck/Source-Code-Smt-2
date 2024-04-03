// Buatlah fungsi saja

func hargaSetelahDiskon(h, d int, m bool) float64 {
	// Mengembalikan harga setelah diberi diskon jika pembeli member
	var total float64
	total = float64(h)
	if m {
		total = float64(h) - (float64(h) * float64(d) / 100)
	}
	return total
}