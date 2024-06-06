// buatlah prosedurnya saja

func gameDadu(d1, d2 int, total *int) {
	/*
	I.S menerima dua bilangan bulat d1 dan d2 sebagai input, dan pointer ke integer total 
	yang diinisialisasi dengan nilai awal 0.
	F.S mengincrement nilai total jika kedua dadu menghasilkan bilangan ganjil.
	*/
	if d1 % 2 != 0 && d2 % 2 != 0{
	    *total += 1
	}
}