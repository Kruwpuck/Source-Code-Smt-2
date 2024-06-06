func sumBulat(bilangan, jumlah, banyak int){
	/*
	IS : Terdefinisi bilangan bulat, yang akan dijumlahkan terus-menerus sampai jumlahnya melewati 100
	FS : Menampilkan hasil penjumlahan bilangan bulat dan berapa kali bilangan bulat diinputkan
	*/
	 banyak = 0
	 jumlah = 0
	 for {
		 fmt.Scan(&bilangan)
		 banyak = banyak + 1
		 jumlah = jumlah + bilangan
		 if jumlah > 100 {
			 break
		 }
	 }
	 fmt.Print(jumlah, banyak)
 }
 