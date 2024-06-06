// Buatlah prosedur saja

func banyakNegatifPositif(b1, b2, b3, b4 int, banyakPositif *int, banyakNegatif *int){
    /*I.S. Terdefinisi bilangan bulat b1,b2,b3,b4, banyakPositif,banyakNegatif
    F.S. Variabel banyakPositif dan banyakNegatif menyimpan banyak bilangan positif 
    dan banyak bilangan negatif
    */
    if b1 > 0 {
		*banyakPositif++
	} else if b1 < 0 {
		*banyakNegatif++
	}

	if b2 > 0 {
		*banyakPositif++
	} else if b2 < 0 {
		*banyakNegatif++
	}

	if b3 > 0 {
		*banyakPositif++
	} else if b3 < 0 {
		*banyakNegatif++
	}

	if b4 > 0 {
		*banyakPositif++
	} else if b4 < 0 {
		*banyakNegatif++
	}
}
