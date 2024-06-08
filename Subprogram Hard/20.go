func cariGo(kar byte) {
	var previousChar byte
	count := 0

	// Scan karakter pertama
	previousChar = kar

	for {
		fmt.Scanf("%c", &kar)
		if kar == '.' {
			break
		}
		if previousChar == 'g' && kar == 'o' {
			count++
		}
		previousChar = kar
	}

	fmt.Println(count)
}