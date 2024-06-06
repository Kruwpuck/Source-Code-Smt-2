func vokalKonsonan(kar byte, jumlahVokal *int, jumlahKonsonan *int) {
	/*  I.S. Terdefinisi kar byte, jumlahVokal int, dan jumlahKonsonan int
		F.S. Jumlah huruf vokal dan konsonan disimpan di jumlahVokal(jumlah huruf vokal) dan jumlahKonsonan(jumlah huruf konsonan) */ 
	
		
		
		for kar != '.'{
					if kar == 'a' || kar == 'i' || kar == 'u' || kar == 'e' || kar == 'o'{
						*jumlahVokal += 1
					}else {
						*jumlahKonsonan += 1
					}
					fmt.Scanf("%c", &kar)
	}
	}