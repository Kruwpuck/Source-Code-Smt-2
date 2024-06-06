func kelipatan(a, b int) {
	/*
	I.S. terdefinisi dua buah bilangan bulat positif a dan b.
	F.S. tercetak "Ya, kelipatan." jika a kelipatan b atau
	"Tidak, bukan kelipatan." jika a bukan kelipatan b.
	*/
   
	if a % b == 0 {
	fmt.Println("Ya,", a, "kelipatan", b)
	} else {
	fmt.Println("Tidak,", a, "bukan kelipatan", b)
	}
	}
   