func cekJenisNode(v byte) {
	if v == 'A' {
	fmt.Print("akar")
	} else if v == 'E' || v == 'B' {
	fmt.Print("verteks dalam")
	} else if v == 'C' || v == 'F' || v == 'G' || v == 'D' {
	fmt.Print("daun")
	}
   }