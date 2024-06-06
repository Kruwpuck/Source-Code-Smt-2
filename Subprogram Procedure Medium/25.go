// buatlah prosedur-prosedurnya saja

func hitungTotalWaktuPutaran(n int, totalWaktu *int) {
	/*
    I.S n menyatakan Jumlah putaran berkuda dan totalWaktu bernilai 0.
    F.S menyimpan total waktu yang ditempuh dan mencetak total catatan waktu berkuda. 
	*/
	var i, putaran int
	for i= 1; i <= n; i++{
	    fmt.Scan(&putaran)
	    *totalWaktu += putaran
		
		}
		fmt.Println(*totalWaktu)
}

func rerataWaktu(n, totalWaktu int) {
	/*
	I.S n menyatakan Jumlah putaran berkuda dan totalWaktu menyatakan total waktu yang ditempuh.
    F.S menghitung dan mencetak rata-rata waktu per putaran. Penulisan rata-rata dibulatkan 2 digit di belakang koma.
	*/
	var rata float64
	rata = float64(totalWaktu)/float64(n)
	fmt.Printf("%.2f", rata)
}



