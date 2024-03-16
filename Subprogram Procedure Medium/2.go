//buatlah procedure saja

func vote(totSiswa, totPemilihA, totPemilihB int, hasil *string) {
	if ((totPemilihA+totPemilihB)*10)/totSiswa >= 6 {
		if totPemilihA > totPemilihB {
			fmt.Print("Kandidat A menang")
		} else {
			fmt.Print("Kandidat B menang")

		}
	} else {
		fmt.Print("Tidak ada pemenang")
	}

}