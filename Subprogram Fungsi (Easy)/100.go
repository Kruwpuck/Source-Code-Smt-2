// Buatlah fungsi saja

func hargaSetelahDiskon(h, d int, m bool, day string) float64 {
	// Mengembalikan harga setelah diberi diskon jika pembeli member
	// dan total belanja lebih dari Rp 100.000
	var total float64
	total = float64(h)
	if m && (day == "Sabtu" || day == "Minggu") {
		total = float64(h) - (float64(h) * float64(d) / 100)
	}
	return total
}