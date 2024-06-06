// buatlah prosedurnya saja
func konsekutifMenaik(b1, b2, b3 int) {
	/*
	I.S. terdefinisi tiga buah bilangan bulat.
	F.S. tercetak "ya" jika konsekutif menaik atau "tidak" jika tidak konsekuti
	*/
	if b2-b1 == 1 && b3-b2 == 1 {
	fmt.Print("ya")
	} else {
	fmt.Print("tidak")
	}
   }
   