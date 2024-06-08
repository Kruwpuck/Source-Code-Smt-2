func tebakAngka(angkaDitebak int) {
	var tebakan int
	percobaan := 0
	digit1 := angkaDitebak / 10
	digit2 := angkaDitebak % 10

	for {
		fmt.Scan(&tebakan)
		percobaan++

		// Check if tebakan is a 2-digit number
		if tebakan >= 10 {
			tebakan1 := tebakan / 10
			tebakan2 := tebakan % 10
			if (tebakan1 == digit1 && tebakan2 == digit2) || (tebakan1 == digit2 && tebakan2 == digit1) {
				break
			}
		} else {
			// Check if tebakan is a 1-digit number
			if tebakan == digit1 || tebakan == digit2 {
				break
			}
		}
	}

	fmt.Println(percobaan)
}
