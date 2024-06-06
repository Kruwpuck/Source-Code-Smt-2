func haveSampoHujan(habisSampo, hujan string) {
	//Menampilkan pergi atau tidak pergi ke minimarket sesuai dengan ketentuan soal
	var hasil string
	if habisSampo == "ya" && hujan == "tidak" {
	hasil = "pergi ke minimarket"
	} else {
	hasil = "tidak pergi ke minimarket"
	}
	fmt.Print(hasil)
   }
   