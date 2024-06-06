func gameSuit(n int, A, B *int) {
	/*
	I.S menerima tiga parameter: bilangan bulat n (jumlah putaran permainan) dan 
	dua pointer ke integer A dan B (skor pemain A dan B). 
	F.S membaca input dari pengguna berupa karakter-karakter yang mewakili pilihan 
	tangan pemain A dan B dalam setiap putaran permainan batu-gunting-kertas. 
	Skor pemain A dan B diupdate berdasarkan aturan permainan.
	*/
	var moves string
	fmt.Scan(&moves)

	for i := 0; i < n; i++ {
		a := moves[2*i]
		b := moves[2*i+1]

		if (a == 'g' && b == 'k') || (a == 'k' && b == 'b') || (a == 'b' && b == 'g') {
			*A += 1
		} else if (b == 'g' && a == 'k') || (b == 'k' && a == 'b') || (b == 'b' && a == 'g') {
			*B += 1
		}
	}
}