func sumKonsonan(jumlah *int) {
	/* I.S. jumlah bernilai 0
	F.S. Nilai jumlah sebanyak huruf a pada rangkaian karakter (perhitugnan berhenti ketika karakter='.')*/
	var kata byte
	
	fmt.Scanf("%c",&kata)
		for kata != '.'{
			if kata == 'a'{
				*jumlah += 1
			}
			fmt.Scanf("%c", &kata)
	}
}