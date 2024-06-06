func haveSampoHujan( habisSampo, hujan string ) {
	/*
	I.S. terdefinisi 2 buah string yang menyatakan status sampo ("ya" berarti kehabisan sampo atau "tidak" bila tidak kehabisan sampo) dan status turun hujan ("ya" berarti hujan turun atau "tidak" berarti tidak turun).
	F.S. tercetak string "pergi ke minimarket" atau "tidak pergi ke minimarket".
	*/
		var hasil string
		if habisSampo == "ya" && hujan == "tidak" {
		hasil = "pergi ke minimarket"
		} else {
		hasil = "tidak pergi ke minimarket"
		}
		fmt.Print(hasil)
	}