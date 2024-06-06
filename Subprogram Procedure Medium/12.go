// Buatlah prosedur saja

func jumlahDeret(n, m int, jumlah *float64) {
	/* I.S Terdefinisi dua bilangan bulat positif yaitu n dan m, dan variabel real jumlah bernilai nol
	   F.S Hasil dari jumlah deret n dan m ditampung oleh variabel jumlah*/
		var i int
		
		for i = n; i <= m; i++{
			*jumlah = *jumlah + (4/float64(i))
		}
	}