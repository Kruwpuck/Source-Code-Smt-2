func hitungJumlah(n int, jumlah *int ) {
	/*  I.S terdefinisi bilangan bulat n yang menyatakan banyaknya bilangan
		F.S jumlah yang menyatakan hasil penjumlahan digit kedua dan digit keempat dari n bilangan */
		var i, bil int
		
		for i = 1; i <= n; i ++{
			fmt.Scan(&bil)
			*jumlah = *jumlah + (bil%100/10) + (bil/1000%10)
		}
	}