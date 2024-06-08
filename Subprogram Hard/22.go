func login(username, password string, totalGagalLogin *int) {
	correctUsername := "admin"
	correctPassword := "admin"

	for username != correctUsername || password != correctPassword {
		*totalGagalLogin++
		fmt.Scan(&username, &password)
	}
}