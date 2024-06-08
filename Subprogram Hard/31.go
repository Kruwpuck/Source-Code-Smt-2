func cetakTerbesar(b1, b2, b3, b4 int) {
	/*
		I.S menerima empat bilangan bulat b1, b2, b3, dan b4.
		F.S mencari dan mencetak bilangan terbesar di antara b1, b2, b3, dan b4.
	*/
	terbesar := b1
	if b2 > terbesar {
		terbesar = b2
	}
	if b3 > terbesar {
		terbesar = b3
	}
	if b4 > terbesar {
		terbesar = b4
	}
	fmt.Println(terbesar)
}
