func denda(hari int) {
	if hari >= 1 && hari <= 5 {
	hari = hari * 1000
	fmt.Println(hari)
	} else if hari >= 6 && hari <= 10 {
	hari = hari * 2000
	fmt.Println(hari)
	} else {
	fmt.Println("cabut keanggotaan")
	}
   }