const nMax int = 100

type dataMobil [nMax]string

func isiData(arrMobil *dataMobil) {
	/*I.S. Terdefinisi nilai array dataMobil yang masih kosong
	  F.S. arrMobil berisi data yang diinputkan*/
	for i := 0; i < nMax; i++ {
		fmt.Scan(&arrMobil[i])
		if arrMobil[i] == "-1" {
			break
		}
	}
}

func hitung(arrMobil dataMobil) string {
	/*Mengembalikan string warna mobil tertentu yang paling banyak
	  melintasi jalan*/
	var n, nMerah, nHitam, nAbu int
	var terbanyak string

	n = len(arrMobil) //menentukan panjang dari arrMobil
	for i := 0; i < n; i++ {
		if arrMobil[i] == "merah" {
			nMerah++
		} else if arrMobil[i] == "hitam" {
			nHitam++
		} else if arrMobil[i] == "abu" {
			nAbu++
		}
	}

	if nMerah > nAbu && nMerah > nHitam {
		terbanyak = "merah"
	} else if nAbu > nMerah && nAbu > nHitam {
		terbanyak = "abu"
	} else {
		terbanyak = "hitam"
	}
	return terbanyak
}