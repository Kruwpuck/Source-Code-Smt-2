func pemenang(p1, p2, p3 string) {
	var hasil string
	if p1 == p2 && p1 == p3 && p2 == p3 {
	hasil = "imbang"
	} else if p1 != p2 && p1 != p3 {
	hasil = "pemain 1 pemenang"
	} else if p2 != p1 && p2 != p3 {
	hasil = "pemain 2 pemenang"
	} else if p3 != p1 && p3 != p2 {
	hasil = "pemain 3 pemenang"
	}
	fmt.Print(hasil)
   }