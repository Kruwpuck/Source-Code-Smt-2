// Buatlah prosedur saja

func game(m1, m2, m3 *int) {
	/* I.S Terdefinisi m1, m2, m3 yang merupakan jumlah kemenangan setiap pemain, saat ini bernilai nol. 
	 F.S Prosedur akan melakukan proses inputan berupa angka pilihan dari ketiga pemain. Inputan berakhir apabila 3 angka tersebut adalah 0 semua. m1, m2, m3 bertambah sesuai kemenangan setiap pemain*/
	 var t1, t2, t3 int// t1 : tebakan pemain 1
	 var p1gen, p1gan, p2gen, p2gan, p3gen, p3gan bool

	 
	 fmt.Scan(&t1, &t2, &t3)//meminta inputan tebakan pemain 1, pemain 2 dan pemain 3
	 for t1 != 0 && t2 != 0 && t3 != 0 {
	 p1gen = t1 % 2 == 0
	 p1gan = t1 % 2 != 0
	 p2gen = t2 % 2 == 0
	 p2gan = t2 %2 != 0
	 p3gen = t3 % 2 == 0
	 p3gan = t3 % 2 != 0
		 if p1gen && p2gan && p3gan {
			 *m1 += 1
		 }else if p1gan && p2gen && p3gen{
			 *m1 += 1 
		 } else if p1gan && p2gen && p3gan {
			*m2 += 1
		 }else if p1gen && p2gan && p3gen{
			 *m2 += 1
		 }else if p1gan && p2gan && p3gen{
			 *m3 += 1
		 }else if p1gen && p2gen && p3gan{
			 *m3 += 1
		 }else if p1gen && p2gen && p3gen{
			*m1 = *m1
			*m2 = *m2
			*m3 = *m3
		 }else if p1gan && p2gan && p3gan{
			*m1 = *m1
			*m2 = *m2
			*m3 = *m3
		 }
		 fmt.Scan(&t1, &t2, &t3)
	 }
	 
 }