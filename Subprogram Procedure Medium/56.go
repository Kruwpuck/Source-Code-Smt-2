func ganjilGenap(bilangan int) {
	/*I.S. Terdefinisi bilangan bulat dengan 4 digit
	F.S.  Menampilkan bilangan 1 apabila digit awal adalah ganjil dan digit akhir
	adalah genap, dan 0 jika tidak memenuhi syarat tersebut
	*/
	var d1, d4 int
	
	d1 = bilangan / 1000
	d4 = bilangan % 10
	
	if d1 % 2 != 0 && d4 % 2 == 0{
	    fmt.Print("1")
	}else{
	    fmt.Print("0")
	}
}