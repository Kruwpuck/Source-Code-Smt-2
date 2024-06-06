func countVowel(jumlah *int) {
	/* I.S. jumlah bernilai 0
	   F.S. jumlah bernilai sebanyak banyaknya huruf 'a' dan 'i' pada rangkaian karakter (perhitungan berhenti ketika karakter='.')*/
		var kata byte
		
		fmt.Scanf("%c",&kata)
			for kata != '.'{
				if kata == 'a'|| kata == 'i'{
					*jumlah += 1
				}
				fmt.Scanf("%c", &kata)
		}
	}