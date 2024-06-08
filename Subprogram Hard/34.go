func dectobin(n int) {
	if n == 0 {
		return
	}

	dectobin(n / 2)

	fmt.Print(n % 2)
}