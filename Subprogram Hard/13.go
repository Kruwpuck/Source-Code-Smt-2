func jumlahKelipatan(bil1, bil2, x int) int {
	// Tukar nilai bil1 dan bil2 jika bil1 lebih besar dari bil2
	if bil1 > bil2 {
		bil1, bil2 = bil2, bil1
	}

	total := 0
	for i := bil1; i <= bil2; i++ {
		if i%x == 0 {
			total += i
		}
	}
	return total
}
