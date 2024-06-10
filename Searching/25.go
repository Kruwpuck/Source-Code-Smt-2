func faktor(n int) {
	/* I.S. terdiri dari sebuah bilangan bulat positif N
	   F.S. terdiri dari kumpulan pasangan nilai bulat d dan boolean s yang dipisahkan oleh newline.
	   Di mana d adalah bilangan yang mungkin menjadi faktor dari N, dan s menyatakan apakah d adalah faktor dari N.
	   Bilangan bulat d dan boolean s dipisah dengan spasi */

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i, true)
		} else {
			fmt.Println(i, false)
		}
	}
}