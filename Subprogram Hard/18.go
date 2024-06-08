func konsekutif(bilangan int) {
	isKonsekutif := true
	previousDigit := -1

	for bilangan > 0 {
		currentDigit := bilangan % 10
		bilangan /= 10

		if previousDigit != -1 && previousDigit-currentDigit != 1 && previousDigit-currentDigit != -1 {
			isKonsekutif = false
			break
		}

		previousDigit = currentDigit
	}

	if isKonsekutif {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}