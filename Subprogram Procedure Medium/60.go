func pointKlub(golMasuk, golKemasukan int) {
    /*IS : Terdefinisi bilangan golMasuk dan golKemasukan sebanyak 38 kali
    FS : Menampilkan jumlah point, jumlah gol memasukkan, jumlah gol kemasukan,
    dan selisih gol dalam 1 musim. 
    */
	var point, totgol, totkemasuk int
	for i := 0; i < 38; i++ {
		fmt.Scanln(&golMasuk,&golKemasukan)
			totgol += golMasuk
			totkemasuk +=golKemasukan
			if golMasuk > golKemasukan{
				point+= 3
			}else if golMasuk == golKemasukan{
				point += 1
			}
			
	}
	fmt.Println(point, totgol,totkemasuk, totgol-totkemasuk)
}
